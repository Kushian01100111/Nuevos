package kata

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	nodes := make([]int, 0)

	node := head

	for node != nil {
		nodes = append(nodes, node.Val)
		node = node.Next
	}

	node = head
	for i := len(nodes); i > 0; i-- {
		node.Val = nodes[i-1]
		node = node.Next
	}

	return head
}
