package kata

func isValid(s string) bool {
	stack := []byte{}

	if len(s) == 1 || len(s) == 0 {
		return false
	}

	i := 0
	for i < len(s) {
		curr := s[i]
		if curr == '(' || curr == '{' || curr == '[' {
			stack = append(stack, curr)
		} else if curr == ')' || curr == '}' || curr == ']' {
			if len(stack) == 0 {
				return false
			}
			if curr == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if curr == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else if curr == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		i++
	}
	return len(stack) == 0
}
