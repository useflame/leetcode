package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testTable := []struct {
		input  []string
		output string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"reflower", "flow", "flight"}, ""},
		{[]string{"a"}, "a"},
	}

	for _, item := range testTable {
		output := LongestCommonPrefix(item.input)
		if output != item.output {
			t.Errorf("Got: %s, want: %s.", output, item.output)
		}
	}
}
