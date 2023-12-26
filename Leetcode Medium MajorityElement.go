package kata

func majorityElement(nums []int) []int {
	amount := float64(len(nums)) / 3.0
	element := map[int]int{}
	result := []int{}
	for _, num := range nums {
		if _, present := element[num]; !present {
			element[num] = 1
		} else {
			element[num]++
		}
	}

	for index, value := range element {
		tempVal := float64(value)
		if tempVal > amount {
			result = append(result, index)
		}
	}

	return result
}
