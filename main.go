package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type config struct {
	DBuser    string
	DBpw      string
	UseDBAuth bool
}

type Puzzle struct {
	Id    string  `json:id`
	Key   uint64  `json:key`
	Value []uint8 `json:value`
}

type Response struct {
	TotalRows uint32   `json:"total_rows,omitempty"`
	Offset    uint8    `json:"offset, omitempty"`
	Rows      []Puzzle `json:"rows"`
}

func db_port() string {
	port := os.Getenv("DB_PORT")
	if port == "" {
		return string(5984)
	}
	return port
}

func db_hostname() string {
	host := os.Getenv("DB_HOSTNAME")
	if host == "" {
		return "localhost"
	}
	return host
}

func admin() string {
	user := os.Getenv("DB_USER")
	if user == "" {
		return "admin"
	}
	return user
}

func pw() string {
	pw := os.Getenv("DB_PW")
	if len(pw) == 0 {
		log.Fatal("unable to retrieve db password (DB_PW not set)")
	}
	return pw
}

func puzzle_count() uint32 {
	req, _ := http.NewRequest("GET", "http://localhost:5984/sudoku/_design/puzzles/_view/completed", nil)
	req.SetBasicAuth(admin(), pw())
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("limit", "1")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("houston, we have a problem: ", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var val Response
	err = decoder.Decode(&val)
	if err != nil {
		log.Fatal(err)
	}

	return val.TotalRows
}

func pick(n uint32) uint32 {
	fmt.Printf("pick a number from 0 to %v\n", n)
	rand.Seed(time.Now().UnixNano())
	p := rand.Int31n(int32(n))
	return uint32(p)
}

func get_puzzle(n uint32) Puzzle {

	req, _ := http.NewRequest("GET", "http://localhost:5984/sudoku/_design/puzzles/_view/completed", nil)
	req.SetBasicAuth(admin(), pw())
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("limit", "1")
	q.Add("skip", string(n))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("houston, we have a problem: ", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var r Response
	err = decoder.Decode(&r)
	if err != nil {
		log.Fatal(err)
	}

	if len(r.Rows) == 0 {
		log.Fatal("unable to retrieve puzzle (rows len == 0)")
	}
	return r.Rows[0]
}

func main() {
	nrow := puzzle_count()
	log.Printf("number of puzzles: %v\n", nrow)
	pick := pick(nrow)
	log.Printf("pick: %v\n", pick)
	puzzle := get_puzzle(pick)

	fmt.Printf("puzzle: %v\n", puzzle)
}
