package kata

import "regexp"

func halvesAreAlike(s string) bool {
	a, b := []byte(s[0:(len(s)/2)]), []byte(s[(len(s)/2):])

	match := regexp.MustCompile("[aeiouAEIOU]")

	return len(match.FindAllIndex(a, -1)) == len(match.FindAllIndex(b, -1))
}
