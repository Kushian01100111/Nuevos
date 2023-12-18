package main

func removeElement(nums []int, val int) int {
	changedNums := []int{}
	countNotVal := 0

	for _, num := range nums {
		if num == val {
			continue
		} else {
			countNotVal++
			changedNums = append(changedNums, num)
		}
	}

	for i, inter := range changedNums {
		nums[i] = inter
	}

	return countNotVal
}
