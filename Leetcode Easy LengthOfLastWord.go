package kata

import "strings"

func lengthOfLastWord(s string) int {
	str := strings.Trim(s, " ")

	i := len(str) - 1
	count := 0
	for i >= 0 {
		if str[i] != ' ' {
			count++
			i--
		} else {
			return count
		}
	}

	return count
}
