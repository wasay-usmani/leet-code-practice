package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced returns true if the binary tree is height-balanced.
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	dif := maxDepth(root.Left) - maxDepth(root.Right)
	if dif <= 1 && dif >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
