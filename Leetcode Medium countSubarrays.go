package kata

import "slices"

///Correct answer

func countSubarrays(nums []int, k int) int64 {
	maxValue := 0
	var maxValueIds []int
	var ans int64

	for i, x := range nums {
		if x > maxValue {
			maxValue, ans, maxValueIds = x, 0, []int{}
		}

		if x == maxValue {
			maxValueIds = append(maxValueIds, i)
		}

		if len(maxValueIds) >= k {
			ans += int64(maxValueIds[len(maxValueIds)-k]) + 1
		}
	}

	return ans
}

//My answer

func CountSubarrays(nums []int, k int) int64 {
	if k > len(nums) {
		return 0
	}

	var count int64
	maxNum := slices.Max(nums)

	diff := k
	left, right := 0, k
	for diff <= len(nums) {
		arr := nums[left:right]
		amount := 0

		if slices.Contains(arr, maxNum) {
			for _, num := range arr {
				if num == maxNum {
					amount++
				}
			}
		}

		if amount >= k {
			count++
		}

		if right >= len(nums) {
			diff++
			left = 0
			right = diff
		} else {
			left++
			right++
		}
	}

	return count
}
