package puzzler

type IntSet struct {
	setMap map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{
		setMap: make(map[int]bool),
	}
}

func (s IntSet) Add(n int) {
	if _, ok := s.setMap[n]; !ok {
		s.setMap[n] = true
	}
}

func (s IntSet) Remove(n int) {
	if _, ok := s.setMap[n]; ok {
		s.setMap[n] = false
	}
	// item not found
}

func (s IntSet) Contains(n int) bool {
	v, ok := s.setMap[n]
	return ok && v
}

func (s IntSet) IsEmpty() bool {
	return s.Size() == 0
}

func (s IntSet) Size() int {
	return len(s.Members())
}

func (s IntSet) Members() []int {
	var r []int
	for k, v := range s.setMap {
		if v {
			r = append(r, k)
		}
	}

	return r
}
