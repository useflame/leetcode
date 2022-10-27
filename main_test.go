// ## exaplme 1
//     Input: s = "III"
//     Output: 3
//     Explanation: III = 3.

// ## Example 2:

//         Input: s = "LVIII"
//         Output: 58
//         Explanation: L = 50, V= 5, III = 3.

// ## Example 3:

//     Input: s = "MCMXCIV"
//     Output: 1994
//     Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
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
