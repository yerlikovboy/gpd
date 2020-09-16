package sudoku

import "gpd/utils"

type Grid [81]uint8

type Board struct {
	OriginID  string
	Timestamp uint64
	Cells     Grid
}

func (b Board) IsSolved() bool {

	for _, v := range b.Cells {
		if v == 0 {
			return false
		}
	}

	return true
}

func (b Board) ClueCount() uint8 {
	var c uint8
	c = 0

	for _, v := range b.Cells {
		if v == 0 {
			c++
		}
	}
	return (81 - c)
}

type Dim [9]uint8

func GetRow(rownum uint8) utils.IntSet {
	s := utils.NewIntSet()
	start_idx := int(rownum) * 9
	for i := 0; i < 9; i++ {
		s.Add(start_idx + i)
	}
	return s
}

func GetColumn(colNum uint8) utils.IntSet {
	s := utils.NewIntSet()
	for i := 0; i < 9; i++ {
		s.Add(int(colNum) + (i * 9))
	}
	return s
}
