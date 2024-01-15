package kata

import "slices"

func findWinners(matches [][]int) [][]int {
	win := map[int]int{}
	lose := map[int]int{}

	i := 0
	for i < len(matches) {
		curr := matches[i]
		if _, present := win[curr[0]]; !present {
			win[curr[0]] = 1
		} else {
			win[curr[0]] += 1
		}

		if _, present := lose[curr[1]]; !present {
			lose[curr[1]] = 1
		} else {
			lose[curr[1]] += 1
		}
		i++
	}

	winners := []int{}
	losers := []int{}

	for i := range win {
		if _, present := lose[i]; !present {
			winners = append(winners, i)
		}
	}
	slices.Sort(winners)

	for i, value := range lose {
		if value == 1 {
			losers = append(losers, i)
		}
	}
	slices.Sort(losers)

	res := [][]int{}
	res = append(res, winners)
	res = append(res, losers)
	return res
}
