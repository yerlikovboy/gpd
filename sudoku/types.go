package sudoku

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
