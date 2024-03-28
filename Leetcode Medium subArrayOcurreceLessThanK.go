package kata

func maxSubarrayLength(nums []int, k int) int {
	count := make(map[int]int)
	maxSize := 0
	left, right := 0, 0

	for right < len(nums) {
		// Expand the window
		count[nums[right]]++

		// Shrink the window if any element occurs more than k times
		for count[nums[right]] > k {
			count[nums[left]]--
			left++
		}

		maxSize = max(maxSize, right-left+1)
		right++
	}

	return maxSize
}
