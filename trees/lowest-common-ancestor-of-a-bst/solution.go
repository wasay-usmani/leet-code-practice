package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor returns the lowest common ancestor of p and q in the BST.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
        return nil
    }
    if p == q || p == root || q == root {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }else{
        if left != nil{
            return left
        }
        return right
    }

}
