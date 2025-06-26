package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder returns the level order traversal of a binary tree's nodes' values.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) != 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			level = append(level, curr.Val)
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		if len(level) != 0 {
			result = append(result, level)
		}
	}

	return result
}
