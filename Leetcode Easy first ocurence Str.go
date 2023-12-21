package kata

import "strings"

func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	}

	i := 0
	expresion := i + len(needle)
	for expresion <= len(haystack) {
		temp := haystack[i:expresion]
		if temp == needle {
			return i
		} else {
			i++
			expresion++
		}
	}
	return i
}
