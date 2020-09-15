package sudoku

import "testing"

func TestNewPicker(t *testing.T) {
	p := NewPickerWithSeed(42)
	if p == nil {
		t.Fail()
	}

	if len(p.rowCount[0].Members()) != 9 {
		t.Fail()
	}
}
