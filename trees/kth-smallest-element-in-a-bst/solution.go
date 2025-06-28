package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest returns the kth smallest element in a BST.
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var count int
	curr := root
	stack := []*TreeNode{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		count += 1
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if count == k {
			return curr.Val
		}

		curr = curr.Right
	}

	return 0
}
