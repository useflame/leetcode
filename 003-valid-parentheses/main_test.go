package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testTable := []struct {
		input  string
		output bool
	}{
		// {"()", true},
		// {"()[]{}", true},
		// {"(]", false},
		// {"([)]", false},
		// {"[(])", false},
		{"()", true},
	}

	for _, item := range testTable {
		output := IsValid(item.input)
		if output != item.output {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}
}
