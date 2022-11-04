package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testTable := []struct {
		input  []int
		output int
	}{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}

	for _, item := range testTable {
		output := removeDuplicates(item.input)

		if output != item.output {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}
}
