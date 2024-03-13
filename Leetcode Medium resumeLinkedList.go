package kata

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	list := &ListNode{}
	mock := list
	arrList := []int{}
	sums := map[int]int{}

	currSum := 0
	i := 0
	for head != nil {
		val := head.Val

		currSum += val
		if index, present := sums[currSum]; present {
			arrList = arrList[:index+1]
			head = head.Next
			i = i - index
			continue
		}
		if currSum == 0 {
			arrList = []int{}
			sums = map[int]int{}
			head = head.Next
			i = 0
			continue
		}
		arrList = append(arrList, val)
		sums[currSum] = i
		head = head.Next
		i++
	}

	for _, val := range arrList {
		mock.Next = &ListNode{Val: val}

		mock = mock.Next
	}

	return list.Next
}

func RemoveZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}

	m := make(map[int]*ListNode)
	m[0] = dummy

	cur := head
	sum := 0
	for cur != nil {
		sum += cur.Val
		m[sum] = cur
		cur = cur.Next
	}

	sum = 0
	cur = dummy
	for cur != nil {
		sum += cur.Val
		cur.Next = m[sum].Next
		cur = cur.Next
	}

	return dummy.Next
}
