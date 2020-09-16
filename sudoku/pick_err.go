package sudoku

// Contigs is a really bad name for a data structure that denotes areas in a row or column where
// there are more than one values next to each other.
// For example, in the following rows, there are two contiguous parts:
//
// 	first Contigs would start at index 1 and end at index 5 and have length of 3
// _  1 2 4 _ _ 7 8 9
type Contigs struct {
	Start  uint8
	End    uint8
	Length uint8
}

func NewContigs(s, e int) Contigs {
	return Contigs{
		Start:  uint8(s),
		End:    uint8(e),
		Length: uint8(e - s),
	}
}

func IsEqual(l, r *Contigs) bool {
	return l.Start == r.Start &&
		l.End == r.End &&
		l.Length == r.Length
}

func Gaps(r Dim) []Contigs {
	// -1 means we currently are not on a contiguous set of clues
	s := -1

	res := make([]Contigs, 0)
	for i, v := range r {
		if v != 0 {
			if s == -1 {
				// found a clue
				s = i
			}
		}

		if v == 0 && s != -1 {
			res = append(res, NewContigs(s, i))
			s = -1
		}
	}
	// if its at the end
	if s != -1 {
		res = append(res, NewContigs(s, 9))
	}
	return res
}
