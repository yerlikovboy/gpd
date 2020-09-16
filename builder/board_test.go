package builder

import "testing"

var rowtests = []struct {
	tc       uint8
	expected []uint8
}{
	{0, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	{1, []uint8{9, 10, 11, 12, 13, 14, 15, 16, 17}},
	{2, []uint8{18, 19, 20, 21, 22, 23, 24, 25, 26}},
	{3, []uint8{27, 28, 29, 30, 31, 32, 33, 34, 35}},
	{4, []uint8{36, 37, 38, 39, 40, 41, 42, 43, 44}},

	{5, []uint8{45, 46, 47, 48, 49, 50, 51, 52, 53}},
	{6, []uint8{54, 55, 56, 57, 58, 59, 60, 61, 62}},
	{7, []uint8{63, 64, 65, 66, 67, 68, 69, 70, 71}},
	{8, []uint8{72, 73, 74, 75, 76, 77, 78, 79, 80}},
}

func TestGetRow(t *testing.T) {

	for _, v := range rowtests {
		actual := getRow(v.tc)

		found := false
		for _, ev := range v.expected {
			for _, v := range actual {
				if ev == v {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("value not found: %v, actual: %v", ev, actual)

			}
		}
	}
}
