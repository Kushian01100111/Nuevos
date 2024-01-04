package kata

import "slices"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	l := len(nums)

	slices.Sort(nums)
	for i := 0; i <= l-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j <= l-2; j++ {
			left, right := j+1, l-1

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}

			}
		}
	}
	return result
}
