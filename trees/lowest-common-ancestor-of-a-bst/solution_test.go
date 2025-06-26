package trees

import "testing"

func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	for i, node := range nodes {
		if node == nil {
			continue
		}
		li, ri := 2*i+1, 2*i+2
		if li < len(vals) {
			node.Left = nodes[li]
		}
		if ri < len(vals) {
			node.Right = nodes[ri]
		}
	}
	return nodes[0]
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		root     []interface{}
		p, q     int
		expected int
	}{
		{
			name:     "Example 1",
			root:     []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5},
			p:        2,
			q:        8,
			expected: 6,
		},
		{
			name:     "Example 2",
			root:     []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5},
			p:        2,
			q:        4,
			expected: 2,
		},
		{
			name:     "Root is LCA",
			root:     []interface{}{2, 1},
			p:        2,
			q:        1,
			expected: 2,
		},
		{
			name:     "Test case",
			root:     []interface{}{3, 1, 4, nil, 2},
			p:        2,
			q:        3,
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.root)
			p := findNode(root, tc.p)
			q := findNode(root, tc.q)
			lca := lowestCommonAncestor(root, p, q)
			if lca == nil || lca.Val != tc.expected {
				t.Errorf("expected %d, got %v", tc.expected, lca)
			}
		})
	}
}
