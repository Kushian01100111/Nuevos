package kata

func climbStairs(n int) int {
	nums := []int{1, 1}
	i := 1
	for i < n {
		temp := nums[i] + nums[i-1]
		nums = append(nums, temp)
		i++
	}

	return nums[len(nums)-1]
}
