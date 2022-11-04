package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {

	testTable := []struct {
		a      []int
		b      int
		output int
	}{
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		},
	}

	for _, i := range testTable {
		output := removeElement(i.a, i.b)
		if output != i.output {
			t.Errorf("Got: %v, want: %v.", output, i.output)
		}
	}
}
