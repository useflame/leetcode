package leetcode

import (
	"sort"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		l1, l2 := len(strs[i]), len(strs[j])
		return l1 < l2
	})

	smallestString := strs[0]

	for i := len(smallestString); i > 0; i-- {
		if isAllHasPrefix(smallestString[:i], strs) {
			return smallestString[:i]
		}
	}

	return ""
}

func isAllHasPrefix(str string, strs []string) bool {
	for _, i := range strs {
		if !strings.HasPrefix(i, str) {
			return false
		}
	}
	return true
}
