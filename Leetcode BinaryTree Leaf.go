package kata

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	end1, end2 := []int{}, []int{}

	leafValue(root1, &end1)
	leafValue(root2, &end2)

	// Compare leaf sequences in the same order
	if len(end1) != len(end2) {
		return false
	}

	for i := 0; i < len(end1); i++ {
		if end1[i] != end2[i] {
			return false
		}
	}

	return true
}

func leafValue(root *TreeNode, end *[]int) int {

	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		*end = append(*end, root.Val)
		return 0
	} else if root.Right == nil {
		return leafValue(root.Left, end)
	} else if root.Left == nil {
		return leafValue(root.Right, end)
	}

	leafValue(root.Left, end)
	return leafValue(root.Right, end)
}
