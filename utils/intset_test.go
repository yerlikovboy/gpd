package utils

import "testing"

func TestNewIntSet(t *testing.T) {

	t_o := NewIntSet()
	if len(t_o.setMap) != 0 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {

	t_o := NewIntSet()

	t_o.Add(1)
	if len(t_o.setMap) != 1 {
		t.Errorf("# of elements in set not 1: %v", len(t_o.setMap))
	}

	if v, ok := t_o.setMap[1]; !ok || !v {
		t.Errorf("incorrect value: %v (or value not found: %v", v, ok)
	}
}

func TestRemove(t *testing.T) {
	t_o := NewIntSet()

	t_o.Add(1)
	if len(t_o.setMap) != 1 {
		t.Errorf("# of elements in set not 1: %v", len(t_o.setMap))
	}

	if v, ok := t_o.setMap[1]; !ok || !v {
		t.Errorf("incorrect value: %v (or value not found: %v", v, ok)
	}

	t_o.Remove(1)
	if v, ok := t_o.setMap[1]; !ok || v {
		t.Errorf("value is still in set: %v (or value not found at all: %v", v, ok)
	}

	if len(t_o.setMap) != 1 {
		t.Errorf("# of elements in set not 1: %v", len(t_o.setMap))
	}
}

func TestContains(t *testing.T) {
	ts := NewIntSet()

	ts.Add(1)
	ts.Add(-1)

	for _, v := range []int{0, 2, -2} {
		if ts.Contains(v) {
			t.Errorf("set contains value %v when it should not", v)
		}
	}

	for _, v := range []int{1, -1} {
		if !ts.Contains(v) {
			t.Errorf("set does not contain value %v when it should ", v)
		}
	}
}

func TestIsEmptyAndSize(t *testing.T) {
	ts := NewIntSet()

	if ts.IsEmpty() == false {
		t.Errorf("isEmpty should be true for newly created set")
	}
	if ts.Size() != 0 {
		t.Errorf("incorrect size (expected: 0, actual: %v)", ts.Size())
	}

	ts.Add(1)
	if ts.IsEmpty() == true {
		t.Errorf("isEmpty should be false for set with one item")
	}
	if ts.Size() != 1 {
		t.Errorf("incorrect size (expected: 1, actual: %v)", ts.Size())
	}
	ts.Remove(1)
	if ts.IsEmpty() == false {
		t.Errorf("isEmpty should be true for set with item added then removed")
	}
	if ts.Size() != 0 {
		t.Errorf("incorrect size (expected: 0, actual: %v)", ts.Size())
	}
}
