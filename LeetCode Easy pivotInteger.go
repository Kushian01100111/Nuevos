package kata

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}

	listNums := make([]int, n)
	num := -1

	for i := 1; i <= n; i++ {
		listNums[i-1] = i
	}

	for i := 1; i < len(listNums); i++ {
		sumLeft := sumArr(listNums[:i+1])
		sumRight := sumArr(listNums[i:])

		if sumLeft == sumRight {
			return listNums[i]
		}
	}

	return num
}

func sumArr(a []int) (sum int) {
	for _, num := range a {
		sum += num
	}
	return
}

func PivotInteger(n int) int {
	listSums := make([]int, n+1)

	for i := 1; i <= n; i++ {
		listSums[i] = i + listSums[i-1]
	}

	for i := 1; i <= n; i++ {
		if listSums[i] == (listSums[n] - listSums[i-1]) {
			return i
		}
	}

	return -1
}
