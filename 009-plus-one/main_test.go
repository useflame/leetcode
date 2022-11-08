package leetcode

import "testing"

func TestPlusOne(t *testing.T) {
	testTable := []struct {
		input  []int
		output []int
	}{
		// {[]int{1, 2, 3}, []int{1, 2, 4}},
		// {[]int{9}, []int{1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}

	for _, item := range testTable {
		output := plusOne(item.input)
		if !isArrayEqual(output, item.output) {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}
}

func isArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
