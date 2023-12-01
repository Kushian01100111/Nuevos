package kata

func isPalindrome(x int) bool {
	initial := x

	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	reverseNum := 0
	remaider := 0
	for {
		if x <= 0 {
			break
		}

		remaider = x % 10
		reverseNum = reverseNum*10 + remaider
		x = int(x / 10)
	}

	if initial == reverseNum {
		return true
	} else {
		return false
	}
}
