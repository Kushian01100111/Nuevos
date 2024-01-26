package kata

func plusOne(digits []int) []int {
	last := len(digits)
	if digits[last-1] != 9 {
		digits[last-1]++
		return digits
	}

	for i := last - 1; i >= 0; i-- {
		temp := digits[i]
		if temp == 9 {
			digits[i] = 0
			continue
		} else {
			digits[i]++
			return digits
		}
	}

	return append([]int{1}, digits...)
}
