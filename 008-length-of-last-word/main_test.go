package leetcode

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testTable := []struct {
		input  string
		output int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, item := range testTable {
		output := lengthOfLastWord(item.input)

		if output != item.output {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}

}
