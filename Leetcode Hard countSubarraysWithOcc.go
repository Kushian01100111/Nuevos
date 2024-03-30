package kata

func subarraysWithAtMostKDistinct(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	occ := make(map[int]int)
	diffInt := 0
	left := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		occ[nums[right]]++
		if occ[nums[right]] == 1 {
			diffInt++
		}

		for diffInt > k {
			occ[nums[left]]--
			if occ[nums[left]] == 0 {
				diffInt--
			}
			left++
		}

		result += right - left + 1
	}

	return result
}

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}
