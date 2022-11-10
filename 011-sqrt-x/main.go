package leetcode

func mySqrt(x int) int {
	i := 1
	for ; i*i <= x; i++ {
		if i*i == x {
			return i
		}
	}

	return i - 1
}
