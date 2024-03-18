package kata

func numSubarraysWithSum(nums []int, goal int) int {
	listSums := make([]int, len(nums)+1)
	amount := 0

	currSum := 0
	for i := 1; i < len(listSums); i++ {
		if nums[i-1] == goal {
			amount++
		}
		currSum += nums[i-1]
		listSums[i] = currSum
	}

	for i := 1; i < len(listSums); i++ {
		for j := i + 1; j < len(listSums); j++ {
			if listSums[j]-listSums[i-1] == goal {
				amount++
			}
		}
	}

	return amount
}
