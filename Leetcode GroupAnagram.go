package kata

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	r := map[string][]string{}

	for _, char := range strs {
		temp := strings.Split(char, "")
		sort.Strings(temp)
		s := strings.Join(temp, "")

		r[s] = append(r[s], char)
	}

	res := [][]string{}

	for _, t := range r {
		res = append(res, t)
	}

	return res
}
