package kata

func findDuplicate(nums []int) int {
	list := make(map[int]int, 0)

	i := 0
	for i < len(nums) {
		list[nums[i]]++
		if val := list[nums[i]]; val > 1 {
			return nums[i]
		}
		i++
	}

	return 0
}
