package sudoku

type Row [9]uint8

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

func Gaps(r Row) []Contigs {

	s := -1

	res := make([]Contigs, 0)
	for i, v := range r {
		if v != 0 {
			if s == -1 {
				// found one
				s = i
			}

		}

		if v == 0 && s != -1 {
			//finished contiguous
			//log.Printf("finish contiguous s: %v, e: %v\n", s, e)
			res = append(res, NewContigs(s, i))
			s = -1
		}
	}
	// if its at the end
	if s != -1 {
		//log.Printf("finish contiguous s: %v, e: %v\n", s, e)
		res = append(res, NewContigs(s, 9))
	}
	return res
}
