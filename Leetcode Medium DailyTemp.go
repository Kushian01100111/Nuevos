package kata

func dailyTemperatures(temps []int) []int {
	recount := make([]int, len(temps))
	stack := make([]int, 0)

	for i, temp := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			recount[index] = i - index
		}
		stack = append(stack, i)
	}
	return recount
}
