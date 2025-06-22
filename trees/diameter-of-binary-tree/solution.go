package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// diameterOfBinaryTree returns the length of the diameter of the binary tree.
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diam := maxDepth(root.Left) + maxDepth(root.Right)
	return max(diam, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
