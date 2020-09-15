package sudoku

import (
	"time"
)

func MakePuzzle(g Board, n_clues uint8) Board {

	picker := NewPicker()

	var i uint8
	i = 0
	n_blanks := 81 - n_clues
	for i < n_blanks {
		p := picker.Pick()

		if g.Cells[p] != 0 {
			g.Cells[p] = 0
			i++
		}
	}

	var grid Grid
	copy(grid[:], g.Cells[0:81])

	return Board{
		Cells:     grid,
		OriginID:  g.OriginID,
		Timestamp: uint64(time.Now().UnixNano()),
	}
}
