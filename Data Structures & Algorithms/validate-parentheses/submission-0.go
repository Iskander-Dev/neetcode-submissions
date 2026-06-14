func isValid(s string) bool {
    stack := []rune{}

	for _, b := range s {
		hasItems := len(stack) > 0
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		} else if hasItems && b == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if hasItems && b == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else if hasItems && b == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
