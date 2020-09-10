package puzzler

import (
	"math/rand"
	"time"
)

type Picker struct {
	// key: # of times picked, values: set of rows
	rowCount map[int]IntSet
}

func initSet() IntSet {
	s := NewIntSet()

	for i := 0; i < 9; i++ {
		s.Add(i)
	}
	return s
}

func NewPicker() Picker {
	return NewPickerWithSeed(time.Now().UnixNano())
}

func NewPickerWithSeed(seed int64) Picker {
	rand.Seed(seed)
	p := Picker{
		rowCount: make(map[int]IntSet),
	}

	p.rowCount[0] = initSet()

	return p
}

func (p Picker) GetN(n int) IntSet {

	if _, ok := p.rowCount[n]; !ok {
		p.rowCount[n] = NewIntSet()
	}

	return p.rowCount[n]
}

func (p Picker) TimesPicked(r int) int {
	if r > 8 {
		return 0
	}

	for i := 0; ; i++ {
		nr := p.rowCount[i]
		if nr.Contains(r) {
			return i
		}
	}
}

//IncRow increments the number of times a row has been picked.
func (p Picker) IncrRow(rownum, curr_n int) {
	p.rowCount[curr_n].Remove(rownum)
	// if there isnt a set of rows for curr_n + 1 # of calls, make one
	if _, ok := p.rowCount[curr_n+1]; !ok {
		p.rowCount[curr_n+1] = NewIntSet()
	}

	p.rowCount[curr_n+1].Add(rownum)
}

func (p Picker) Dump() {
	count := 0
	for i := 0; ; i++ {
		r := p.rowCount[i]
		count += r.Size()
		if count >= 9 {
			break
		}
	}
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
