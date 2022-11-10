package leetcode

import "testing"

func TestMySqrt(t *testing.T) {
	testTable := []struct {
		input  int
		output int
	}{
		{4, 2},
		{8, 2},
	}

	for _, item := range testTable {
		output := mySqrt(item.input)

		if output != item.output {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}
}
