package kata

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	num := recurse(root, &sum, low, high)

	return sum - num
}

func recurse(root *TreeNode, sum *int, low int, high int) int {

	if root == nil {
		return 0
	} else if root.Right == nil && root.Left == nil {
		if root.Val <= high && root.Val >= low {
			*sum += root.Val
		}
		return 0
	} else if root.Right == nil {
		if root.Val <= high && root.Val >= low {
			*sum += root.Val
		}
		return recurse(root.Left, sum, low, high)
	} else if root.Left == nil {
		if root.Val <= high && root.Val >= low {
			*sum += root.Val
		}
		return recurse(root.Right, sum, low, high)
	}

	if root.Val <= high && root.Val >= low {
		*sum += root.Val
	}
	recurse(root.Left, sum, low, high)
	return recurse(root.Right, sum, low, high)
}
