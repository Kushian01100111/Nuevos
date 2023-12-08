package kata

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	x := 1
	if s[0] == '-' {
		x = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	result := 0
	for _, current := range s {
		if !unicode.IsDigit(current) {
			break
		}

		curr := int(current - '0')
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && curr >= 7) {
			return math.MaxInt32
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && curr >= 8) {
			return math.MinInt32
		}
		result = result*10 + curr*x
	}
	return result
}
