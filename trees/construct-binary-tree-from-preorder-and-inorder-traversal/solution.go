package main

import "slices"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree constructs a binary tree from preorder and inorder traversal arrays
// preorder: root -> left subtree -> right subtree
// inorder: left subtree -> root -> right subtree
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := slices.Index(inorder, preorder[0])
	lpre, rpre := preorder[1:i+1], preorder[i+1:]
	lin, rin := inorder[:i], inorder[i+1:]

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(lpre, lin),
		Right: buildTree(rpre, rin),
	}
}
