package kata

func buyChoco(prices []int, money int) int {
	amounts := map[int]int{}

	for i := 0; i < len(prices); i++ {
		currentElm := prices[i]
		for x, nums := range prices {
			value := currentElm + nums
			temp := money - value
			if i == x {
				continue
			} else if _, present := amounts[value]; !present && temp > -1 {
				amounts[value] = money - value
			}
		}
	}

	if len(amounts) == 0 {
		return money
	}
	smallestChocos := 0
	for _, value := range amounts {
		if value >= 0 && value > smallestChocos {
			smallestChocos = value
		}
	}

	return smallestChocos
}
