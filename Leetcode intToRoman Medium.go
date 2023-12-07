package kata

func intToRoman(n int) string {
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	char := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	result := ""
	i := 12

	for true {
		temp := n / nums[i]
		n %= nums[i]
		for temp > 0 {
			result += char[i]
			temp--
		}
		i--
		if n < 1 {
			break
		}
	}
	return result
}
