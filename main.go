package leetcode

func RomanToInt(s string) int {
	result := 0
	numMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		s1 := numMap[rune(s[i])]

		if i+1 < len(s) {
			s2 := numMap[rune(s[i+1])]

			if s1 >= s2 {
				result += s1
			} else {
				result += s2 - s1
				i++
			}
		} else {
			result += s1
		}
	}

	return result
}
