package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverts a binary tree and returns its root.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		root.Left, root.Right = root.Right, root.Left
	} else {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}

	return root
}
