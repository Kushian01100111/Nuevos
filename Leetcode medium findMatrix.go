package kata

import "slices"

func findMatrix(nums []int) [][]int {
	result := [][]int{}

	for len(nums) > 0 {
		temp := []int{}

		for i := 0; i < len(nums); i++ {
			expression := slices.Contains(temp, nums[i])
			if !expression {
				temp = append(temp, nums[i])
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
		result = append(result, temp)
	}

	return result
}
