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
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp
	} else {
		tmp := root.Left
		root.Left = invertTree(root.Right)
		root.Right = invertTree(tmp)
	}

	return root
}
