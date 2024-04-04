package kata

func maxDepth(s string) int {
	stack := []byte{}

	longest := 0
	for _, char := range s {
		if char == '(' {
			stack = append(stack, '(')
			if len(stack) > longest {
				longest = len(stack)
			}
		} else if char == ')' {
			stack = stack[:len(stack)-1]
		}
	}

	return longest
}
