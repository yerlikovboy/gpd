package sudoku

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    Row
	expected []Contigs
}

func TestGaps(t *testing.T) {

	tc := []TestCase{
		TestCase{
			input: Row{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []Contigs{
				Contigs{Start: 0, End: 9, Length: 9},
			},
		},
		TestCase{
			input: Row{1, 2, 3, 0, 5, 6, 7, 8, 9},
			expected: []Contigs{
				Contigs{Start: 0, End: 3, Length: 3},
				Contigs{Start: 4, End: 9, Length: 5},
			},
		},
		TestCase{
			input: Row{0, 2, 3, 0, 5, 6, 7, 8, 9},
			expected: []Contigs{
				Contigs{Start: 1, End: 3, Length: 2},
				Contigs{Start: 4, End: 9, Length: 5},
			},
		},
		TestCase{
			input: Row{0, 2, 3, 0, 5, 6, 7, 8, 0},
			expected: []Contigs{
				Contigs{Start: 1, End: 3, Length: 2},
				Contigs{Start: 4, End: 8, Length: 4},
			},
		},
		TestCase{
			input: Row{0, 2, 3, 0, 5, 6, 7, 0, 0},
			expected: []Contigs{
				Contigs{Start: 1, End: 3, Length: 2},
				Contigs{Start: 4, End: 7, Length: 3},
			},
		},
		TestCase{
			input: Row{0, 2, 3, 4, 0, 0, 0, 0, 0},
			expected: []Contigs{
				Contigs{Start: 1, End: 4, Length: 3},
			},
		},
		TestCase{
			input: Row{0, 2, 3, 0, 5, 0, 0, 8, 0},
			expected: []Contigs{
				Contigs{Start: 1, End: 3, Length: 2},
				Contigs{Start: 4, End: 5, Length: 1},
				Contigs{Start: 7, End: 8, Length: 1},
			},
		},
	}

	for _, v := range tc {
		actual := Gaps(v.input)

		if len(actual) != len(v.expected) {
			fmt.Println(actual)
			t.Errorf("incorrect length: expect: %v, actual: %v", len(v.expected), len(actual))
		}

		for _, ec := range v.expected {

			found := false
			for _, ac := range actual {
				if IsEqual(&ac, &ec) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected object not found: expected: %v, actual: %v", ec, actual)
			}
		}

	}

}
