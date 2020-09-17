package builder

type Grid [81]uint8

type Board struct {
	rowIdx map[uint8][9]uint8
	colIdx map[uint8][9]uint8
	grid   Grid
}

func NewBoard(g Grid) Board {

	rowMap := make(map[uint8][9]uint8, 9)
	colMap := make(map[uint8][9]uint8, 9)

	var i uint8
	for i = 0; i < 9; i++ {
		rowMap[i] = getRow(i)
		colMap[i] = getColumn(i)
	}

	return Board{
		rowIdx: rowMap,
		colIdx: colMap,
		grid:   g,
	}
}

func (b Board) NumClues() uint8 {
	var c uint8
	c = 0
	for _, v := range b.grid {
		if v != 0 {
			c++
		}
	}
	return c
}

func (b Board) IsSolved() bool {
	for _, v := range b.grid {
		if v == 0 {
			return false
		}
	}
	return true
}

func (b Board) RowIdx(rowNum uint8) []uint8 {
	rv, _ := b.rowIdx[rowNum]
	return rv[:]
}

func (b Board) ColIdx(colNum uint8) []uint8 {
	rv, _ := b.rowIdx[colNum]
	return rv[:]
}

func getRow(rownum uint8) [9]uint8 {
	var s [9]uint8
	start_idx := rownum * 9
	var i uint8
	for i = 0; i < 9; i++ {
		s[i] = start_idx + i
	}
	return s
}

func getColumn(colNum uint8) [9]uint8 {
	var s [9]uint8
	for i := 0; i < 9; i++ {
		s[i] = colNum + uint8(i*9)
	}
	return s
}
