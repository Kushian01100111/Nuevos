package kata

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	pali := []int{}

	for head != nil {
		pali = append(pali, head.Val)
		head = head.Next
	}

	if len(pali) == 1 || len(pali) == 2 && pali[0] == pali[1] {
		return true
	}

	i := len(pali) / 2
	j := (len(pali) / 2) - 1

	if len(pali)%2 != 0 {
		i++
	}

	for i < len(pali) && j >= 0 {
		if pali[i] != pali[j] {
			return false
		}
		i++
		j--
	}

	return true
}
