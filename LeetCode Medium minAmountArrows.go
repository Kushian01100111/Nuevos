package kata

import "sort"

func findMinArrowShots(points [][]int) int {
	// exploted := make(map[int]int, 0)

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	amount := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			amount++
			end = points[i][1]
		}
	}

	return amount
}
