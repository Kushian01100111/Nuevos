package kata

func removeDuplicates(nums []int) int {
	mapOfNums := map[int]int{}
	result := []int{}
	count := 0

	for _, val := range nums {
		if _, present := mapOfNums[val]; !present {
			result = append(result, val)
			mapOfNums[val] = 1
			count++
		}
	}

	for i, val := range result {
		nums[i] = val
	}
	return count
}
