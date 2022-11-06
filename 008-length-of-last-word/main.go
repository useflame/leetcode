package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	wordsSlice := strings.Split(strings.Trim(s, " "), " ")

	return len(wordsSlice[len(wordsSlice)-1])
}
