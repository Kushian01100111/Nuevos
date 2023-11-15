// Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

// For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

// The input will be a lowercase string with no spaces.

// Good luck!

// If you like this Kata, please try:

package kata

import "strings"

func main(str string) (evenAndOdd []string) {
	evenAndOdd = []string{"", ""}

	for index, char := range str {
		if index%2 == 0 {
			evenAndOdd[0] += strings.ToUpper(string(char))
			evenAndOdd[1] += string(char)
		} else {
			evenAndOdd[0] += string(char)
			evenAndOdd[1] += strings.ToUpper(string(char))
		}
	}
	return
}
