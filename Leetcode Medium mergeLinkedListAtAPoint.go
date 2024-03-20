package kata

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	i := 0
	for list1 != nil && list2 != nil {
		if i == a {
			curr.Next = &ListNode{Val: list2.Val}
			curr = curr.Next
			if list2.Next != nil {
				list2 = list2.Next
			} else {
				i++
				list1 = list1.Next
			}
		} else if i < b+1 && i > a {
			list1 = list1.Next
			i++
		} else {
			curr.Next = &ListNode{Val: list1.Val}
			curr = curr.Next

			i++
			list1 = list1.Next
		}
	}

	return dummy.Next
}
