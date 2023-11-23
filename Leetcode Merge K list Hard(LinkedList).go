package kata

func mergeKLists(lists []*ListNode) *ListNode {
	mergeList := &ListNode{}
	current := mergeList
	finish, fi := 1, 0

	for fi < finish {
		currentLowest, i := 10000000, 0
		change := 0

		for item, list := range lists {
			if list != nil {
				if currentLowest > list.Val {
					currentLowest, i = list.Val, item
					change += 1
				}
			}
		}
		if change > 0 {
			current.Next = &ListNode{Val: currentLowest}
			current = current.Next
			lists[i] = lists[i].Next
		} else {
			fi = 1
		}
	}

	return mergeList.Next
}
