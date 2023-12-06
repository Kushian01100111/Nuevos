package kata

func romanToInt(s string) int {
	nums := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	dobleNums := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	sum := 0

	for i := 0; i < len(s); i++ {
		x := string(s[i])
		if i+1 < len(s) {
			y := string(s[i+1])
			xy := x + y
			if value, present := dobleNums[xy]; present {
				sum += value
				i++
			} else {
				sum += nums[x]
			}
		} else {
			sum += nums[x]
		}
	}
	return sum
}
