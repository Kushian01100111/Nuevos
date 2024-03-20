package kata

import "sort"

func leastInterval(tasks []byte, n int) int {
	elem := map[byte]int{}

	for _, l := range tasks {
		elem[l]++
	}

	frequencies := make([]int, 0, len(elem))
	for _, freq := range elem {
		frequencies = append(frequencies, freq)
	}

	sort.Slice(frequencies, func(a, b int) bool {
		return frequencies[a] > frequencies[b]
	})

	maxNum := frequencies[0]
	i := 0
	for _, n := range frequencies {
		if n == maxNum {
			i++
		} else {
			break
		}
	}

	secundPossi := (maxNum-1)*(n+1) + i

	return max(secundPossi, len(tasks))
}


