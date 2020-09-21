package sudoku

import (
	"log"
	"math/rand"
	"time"
)

// ClueGroup is a data structure that denotes areas in a row or column where
// there are more than one values next to each other.
// For example, in the following rows, there are two contiguous parts:
//
// 	first ClueGroup would start at index 1 and end at index 5 and have length of 3
// _  1 2 4 _ _ 7 8 9
type ClueGroup struct {
	Start  uint8
	Length uint8
}

func NewClueGroup(s, l int) ClueGroup {
	return ClueGroup{
		Start:  uint8(s),
		Length: uint8(l),
	}
}

func IsEqual(l, r *ClueGroup) bool {
	return l.Start == r.Start &&
		l.Length == r.Length
}

func Find(r []uint8) []ClueGroup {
	// -1 means we currently are not on a contiguous set of clues
	s := -1

	var res []ClueGroup
	for i, v := range r {
		if v != 0 {
			if s == -1 {
				// found a clue
				s = i
			}
		}

		if v == 0 && s != -1 {
			res = append(res, NewClueGroup(s, i-s))
			s = -1
		}
	}
	// if its at the end
	if s != -1 {
		res = append(res, NewClueGroup(s, len(r)-s))
	}
	return res
}

func col_groups(g Grid) []ClueGroup {

	var res []ClueGroup

	for c := 0; c < 9; c++ {
		s := -1
		for i := 0; i < 9; i++ {
			// index equals
			idx := i*9 + c
			v := g[idx]
			if v != 0 {
				// found a new clue
				if s == -1 {
					s = idx
				}
			} else if v == 0 && s != -1 {
				// end of clue group
				len := (idx - s) / 9
				res = append(res, NewClueGroup(s, len))
				s = -1
			}
		}
		if s != -1 {
			len := (81 + c - s) / 9
			res = append(res, NewClueGroup(s, len))
		}
	}

	return res
}

func candidates(g Grid) []uint8 {

	var r []uint8

	cgs := Find(g[:])

	log.Printf("cgs: %v", cgs)
	for _, cg := range cgs {

		multiplier := cg.Length
		if multiplier > 9 {
			multiplier = 9
		}
		var i uint8

		for i = 0; i < multiplier; i++ {
			// need to make a slice of numbers which holds the index
			tmp := make([]uint8, cg.Length)
			var j uint8
			for j = 0; j < cg.Length; j++ {
				tmp[j] = cg.Start + j
			}
			r = append(r, tmp[:]...)
		}
	}

	return r
}

func Make(s Board, num_clues uint8) Board {

	log.Printf("making puzzle with %v clues", num_clues)
	rand.Seed(time.Now().UnixNano())
	var lc Grid
	cells := s.Cells()
	copy(lc[:], cells[:])

	var i uint8
	num_zeros := 81 - num_clues
	for i = 0; i < num_zeros; i++ {
		c := candidates(lc)
		pick_idx := rand.Intn(len(c))
		pick := c[pick_idx]
		log.Printf("pick_idx: %v, pick: %v, len(c): %v", pick_idx, pick, len(c))
		lc[pick] = 0
	}

	return NewBoard(lc).
		WithDerivedFromID(s.ID).
		WithCreatedTS(uint64(time.
			Now().
			UnixNano()))
}
