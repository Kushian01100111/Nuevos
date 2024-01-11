package kata

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	recount := 0
	// elm:= map[int]int{}

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			recount++
			i++
		}
		j++
	}

	return recount
}
