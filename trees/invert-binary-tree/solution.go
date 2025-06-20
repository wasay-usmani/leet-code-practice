package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InvertTree inverts a binary tree and returns its root.
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}
