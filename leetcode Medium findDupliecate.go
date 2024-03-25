package kata

func findDuplicates(nums []int) []int {
	list := make(map[int]int, 0)
	result := make([]int, 0)

	for _, num := range nums {
		list[num]++
		if val := list[num]; val > 1 {
			result = append(result, num)
		}
	}

	return result
}
