package leetcode

func addBinary(a string, b string) string {
	result := ""
	next := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 {
		n := next
		if i >= 0 && a[i] == '1' {
			n += 1
		}

		if j >= 0 && b[j] == '1' {
			n += 1
		}

		if n%2 == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}

		if n < 2 {
			next = 0
		} else {
			next = 1
		}
		i--
		j--
	}

	if next != 0 {
		result = "1" + result
	}

	return result
}
