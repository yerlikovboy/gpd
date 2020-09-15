package sudoku

import (
	"gpd/utils"
	"math/rand"
	"time"
)

type Picker struct {
	// key: # of times picked, values: set of rows
	rowCount map[int]utils.IntSet
}

func initSet() utils.IntSet {
	s := utils.NewIntSet()

	for i := 0; i < 9; i++ {
		s.Add(i)
	}
	return s
}

func NewPicker() Picker {
	return *NewPickerWithSeed(time.Now().UnixNano())
}

func NewPickerWithSeed(seed int64) *Picker {
	rand.Seed(seed)
	p := Picker{
		rowCount: make(map[int]utils.IntSet),
	}

	p.rowCount[0] = initSet()
	return &p
}

func (p Picker) GetN(n int) utils.IntSet {

	if _, ok := p.rowCount[n]; !ok {
		p.rowCount[n] = utils.NewIntSet()
	}

	return p.rowCount[n]
}

//IncRow increments the number of times a row has been picked.
func (p Picker) IncrRow(rownum, curr_n int) {
	p.rowCount[curr_n].Remove(rownum)
	// if there isnt a set of rows for curr_n + 1 # of calls, make one
	if _, ok := p.rowCount[curr_n+1]; !ok {
		p.rowCount[curr_n+1] = utils.NewIntSet()
	}

	p.rowCount[curr_n+1].Add(rownum)
}

func (p Picker) Pick() int {

	// pick a row ...
	var rowPick int
	for n := 0; ; n++ {

		rows := p.GetN(n)
		members := rows.Members()

		if !rows.IsEmpty() {
			pick := rand.Intn(len(members))
			rowPick = members[pick]

			p.IncrRow(rowPick, n)

			break
		}

	}

	return (rowPick * 9) + rand.Intn(9)
}
