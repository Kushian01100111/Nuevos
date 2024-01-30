package kata

import (
	"sort"
	"strconv"
)

func evalRPN(tokens []string) int {
	val := 0
	arr := []string{"*", "+", "-", "/"}
	nums := []int{}

	i := 0
	for i < len(tokens) {
		temp := tokens[i]
		s := sort.SearchStrings(arr, temp)
		if s < len(arr) && arr[s] == temp {
			se := nums[len(nums)-1]
			nums = append(nums[:len(nums)-1], nums[len(nums):]...)
			fi := nums[len(nums)-1]
			nums = append(nums[:len(nums)-1], nums[len(nums):]...)

			switch temp {
			case "+":
				val = fi + se
			case "-":
				val = fi - se
			case "/":
				val = fi / se
			case "*":
				val = fi * se
			}
			nums = append(nums, val)
			i++
		} else if len(tokens) == 1 {
			v, _ := strconv.Atoi(temp)
			val = v
			i++
		} else {
			v, _ := strconv.Atoi(temp)
			nums = append(nums, v)
			i++
		}
	}
	return val
}
