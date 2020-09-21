package sudoku

import (
	"time"
)

func MakePuzzle(g Board, n_clues uint8) Board {

	picker := NewPicker()

	var i uint8
	i = 0
	n_blanks := 81 - n_clues
	var grid Grid
	s := g.Cells()
	copy(grid[:], s[:])

	for i < n_blanks {
		p := picker.Pick()

		if grid[p] != 0 {
			grid[p] = 0
			i++
		}
	}

	return NewBoard(grid).
		WithDerivedFromID(g.ID).
		WithCreatedTS(uint64(time.Now().UnixNano()))
}
