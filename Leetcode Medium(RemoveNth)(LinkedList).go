// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

// Example 2:

// Input: head = [1], n = 1
// Output: []

// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

//     The number of nodes in the list is sz.
//     1 <= sz <= 30
//     0 <= Node.val <= 100
//     1 <= n <= sz

// Follow up: Could you do this in one pass?

package kata

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	currentDummy := dummy
	list := make([]int, 1)

	for head != nil {
		if head != nil {
			list = append(list, head.Val)
			head = head.Next
		}
	}

	list = append(list[:(len(list)-n)], list[(len(list)-(n-1)):]...)

	for item, value := range list {
		if item == 0 {
			continue
		}
		currentDummy.Next = &ListNode{Val: value}
		currentDummy = currentDummy.Next
	}

	return dummy.Next
}
