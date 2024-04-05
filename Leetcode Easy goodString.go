package kata

func makeGood(s string) string {
	result := make([]rune, 0, len(s))

	for _, char := range s {
		if len(result) > 0 && isUpperOrLower(result[len(result)-1], char) {
			result = result[:len(result)-1]
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func abs(x rune) rune {
	if x < 0 {
		return -x
	}
	return x
}

func isUpperOrLower(x, y rune) bool {
	return abs(x-y) == 32
}
