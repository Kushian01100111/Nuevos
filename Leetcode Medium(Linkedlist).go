package kata

func lengthOfLongestSubstring(s string) int {
	charPlace := map[rune]int{}
	maxLength := 0
	start := 0

	for i, char := range s {
		if index, found := charPlace[char]; found && index >= start {
			start = index + 1
		}
		charPlace[char] = i

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}
