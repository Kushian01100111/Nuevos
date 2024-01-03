package kata

import "slices"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	slices.Sort(nums)
	dummy := &ListNode{}
	current := dummy
	for i := range nums {
		dummy.Next = &ListNode{Val: nums[i]}
		dummy = dummy.Next
	}

	return current.Next
}
