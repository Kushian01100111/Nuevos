package kata

func firstMissingPositive(nums []int) int {
	list := make(map[int]int, 0)

	i := 1
	for _, num := range nums {
		list[num]++
		if num == i {
			for {
				i++
				if _, present := list[i]; !present {
					break
				}
			}
		}
	}
	return i
}
