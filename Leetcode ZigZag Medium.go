package kata

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	row := make([]string, numRows)

	index, step := 0, 1

	for _, value := range s {
		row[index] += string(value)

		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}

		index += step
	}

	result := strings.Join(row, "")

	return result
}
