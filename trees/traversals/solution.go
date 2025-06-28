package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal returns the inorder traversal of the BST.
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	curr := root
	stack := []*TreeNode{}
	visited := []int{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited = append(visited, curr.Val)
		curr = curr.Right
	}

	return visited
}

// PreorderTraversal returns the preorder traversal of the BST.
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	visited := []int{root.Val}
	visited = append(visited, PreorderTraversal(root.Left)...)
	visited = append(visited, PreorderTraversal(root.Right)...)
	return visited
}

// PostorderTraversal returns the postorder traversal of the BST.
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	visited := []int{}
	visited = append(visited, PostorderTraversal(root.Left)...)
	visited = append(visited, PostorderTraversal(root.Right)...)
	visited = append(visited, root.Val)
	return visited
}

// LevelOrderTraversal returns the level order (BFS) traversal of the BST.
func LevelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	visited := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited = append(visited, curr.Val)

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return visited
}
