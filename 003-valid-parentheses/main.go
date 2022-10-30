package leetcode

func IsValid(s string) bool {
	var arr []rune

	for _, i := range s {
		if i == '(' || i == '{' || i == '[' {
			arr = append(arr, i)
			continue
		} else if len(arr) == 0 {
			return false
		}

		top := arr[len(arr)-1]
		arr = arr[:len(arr)-1]

		if top == '(' && i != ')' {
			return false
		} else if top == '[' && i != ']' {
			return false
		} else if top == '{' && i != '}' {
			return false
		}

	}
	return len(arr) == 0
}
