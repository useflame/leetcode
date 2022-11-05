package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	testTable := []struct {
		firstArgument  []int
		secondArgument int
		output         int
	}{}

	for _, item := range testTable {
		output := searchInsert(item.firstArgument, item.secondArgument)
		if item.output != output {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}
}
