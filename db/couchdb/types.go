package couchdb

import (
	"gpd/sudoku"
)

type response struct {
	TotalRows uint32 `json:"total_rows,omitempty"`
	Rows      []grid `json:"rows"`
}

type grid struct {
	ID        string  `json:"id"`
	Timestamp uint64  `json:"key"`
	Value     []uint8 `json:"value"`
}

type Puzzle struct {
	ID              string      `json:"_id,omitempty"`
	NumClues        uint8       `json:"n_clues"`
	Cells           sudoku.Grid `json:"grid"`
	SolutionID      string      `json:"solution_id"`
	GeneratedMillis uint64      `json:"generated_millis"`
}

func FromBoard(b sudoku.Board) Puzzle {
	return Puzzle{
		NumClues:        b.NumClues(),
		Cells:           b.Cells(),
		SolutionID:      b.DerivedFromID(),
		GeneratedMillis: b.CreatedTS(),
	}
}
