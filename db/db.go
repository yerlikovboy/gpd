package db

import (
	"gpd/sudoku"
)

type SudokuDB interface {
	Solution() sudoku.Board
	StorePuzzle(s sudoku.Board)
}
