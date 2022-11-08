package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; ; {
		if digits[i] == 9 {
			digits[i] = 0
			if i-1 < 0 {
				digits = append([]int{1}, digits...)

				return digits
			} else {
				digits[i] = 0
				i--
			}
		} else {
			digits[i]++
			return digits
		}
	}
}
