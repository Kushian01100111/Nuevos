package kata

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	index := 0

	for head != nil {
		if head.Val == current.Val && index != 0 {
			head = head.Next
		} else {
			current.Next = &ListNode{Val: head.Val}
			current = current.Next
			head = head.Next
			index++
		}
	}
	return dummy.Next
}
