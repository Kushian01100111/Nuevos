package kata

import "slices"

func divideArray(nums []int, k int) [][]int {
	arr := make([][]int, 0)
	l := len(nums)

	slices.Sort(nums)

	i := 0
	for i < l {
		iArr := []int{nums[i]}
		j := i + 1
		for i < l && j < l {
			diff := abs(nums[i] - nums[j])

			if diff <= k && len(iArr) == 1 {
				iArr = append(iArr, nums[j])
				j++
			} else if diff <= k && len(iArr) < 3 && (abs(iArr[1]-nums[j])) <= k {
				iArr = append(iArr, nums[j])
				i = j
				j = l
				arr = append(arr, iArr)
			} else {
				j++
			}
		}
		i++
	}

	if len(arr) != (l / 3) {
		return [][]int{}
	}

	return arr
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
