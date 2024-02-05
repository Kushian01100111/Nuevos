package kata

import "strings"

func firstUniqChar(s string) int {
	res := -1

	for i := 0; i < len(s); i++ {
		l := string(s[i])
		temp := s[:i] + s[i+1:]
		if !strings.Contains(temp, l) {
			res = i
			return res
		}
	}

	return res
}
