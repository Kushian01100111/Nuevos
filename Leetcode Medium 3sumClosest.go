package kata

import (
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	result := 0
	n := len(nums)

	// Sort the input array
	slices.Sort(nums)
	smallestDiff := math.MaxInt32 / 2

	for i := 0; i < n-2; i++ {
		// Skip duplicate elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := abs(sum - target)

			if diff == 0 {
				return sum
			}

			if diff < smallestDiff {
				smallestDiff = diff
				result = sum
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
