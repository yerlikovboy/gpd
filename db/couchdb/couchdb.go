package couchdb

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"gpd/db"
	"gpd/sudoku"
)

type CouchSudokuDB struct {
	clnt *http.Client
	cfg  Config
}

func New() db.SudokuDB {
	return CouchSudokuDB{
		clnt: &http.Client{},
		cfg:  DefaultConfig(),
	}
}

// WithSeed is for testing
func WithSeed(seed int64) db.SudokuDB {
	rand.Seed(seed)
	return New()
}

func (s CouchSudokuDB) puzzle_count() uint32 {

	req, _ := http.NewRequest("GET", "http://hostname:5984/grids/_design/puzzles/_view/completed", nil)
	s.cfg.SetupRequest(req)

	// the way query is set is so lame ...
	q := req.URL.Query()
	q.Add("limit", "1")
	req.URL.RawQuery = q.Encode()

	resp, err := s.clnt.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("houston, we have a problem: ", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var val response
	err = decoder.Decode(&val)
	if err != nil {
		log.Fatal(err)
	}

	return val.TotalRows
}

func (s CouchSudokuDB) nth_grid(n uint32) grid {

	// log.Printf("pick #%v from view", n)

	req, _ := http.NewRequest("GET", "http://localhost:5984/grids/_design/puzzles/_view/completed?limit=1", nil)
	s.cfg.SetupRequest(req)

	// the way query is set is bullshit!
	qry := req.URL.Query()
	qry.Add("limit", "1")
	qry.Add("skip", fmt.Sprint(n))
	req.URL.RawQuery = qry.Encode()

	// log.Printf("request: %v", req)
	resp, err := s.clnt.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("houston, we have a problem: ", resp.StatusCode)

	}

	decoder := json.NewDecoder(resp.Body)
	var r response
	err = decoder.Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("response: %v", r)
	if len(r.Rows) == 0 {
		log.Fatal("unable to retrieve puzzle (rows len == 0)")
	}
	return r.Rows[0]
}

func (s CouchSudokuDB) Solution() sudoku.Board {

	rowCount := s.puzzle_count()
	pick := uint32(rand.Int31n(int32(rowCount)))
	grid := s.nth_grid(pick)
	var c sudoku.Grid
	copy(c[:], grid.Value[0:81])
	return sudoku.Board{
		OriginID:  grid.ID,
		Timestamp: grid.Timestamp,
		Cells:     c,
	}
}

func (s CouchSudokuDB) StorePuzzle(b sudoku.Board) {
	// for now, just print it to console
	p := FromBoard(b)
	raw, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(raw))
}
