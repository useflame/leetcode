package leetcode

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testTable := []struct {
		input  string
		output int
	}{
		{"I", 1},
		{"III", 3},
		{"XLIX", 49},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
	}

	for _, item := range testTable {
		output := RomanToInt(item.input)
		if output != item.output {
			t.Errorf("got :%d; want: %d", output, item.output)
		}
	}
}
