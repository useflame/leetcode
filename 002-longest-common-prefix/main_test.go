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
		// {[]string{"racecar", "racecar", "racecar"}, "racecar"},
		// {[]string{"racecar", "acecar", "acecar"}, "acecar"},
		{[]string{"reflower", "flow", "flight"}, ""},
		{[]string{"a"}, "a"},
	}

	for _, item := range testTable {
		output := LongestCommonPrefix(item.input)
		if output != item.output {
			t.Errorf("got :%s; want: %s", output, item.output)
		}
	}
}
