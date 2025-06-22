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

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name     string
		root     []interface{}
		subRoot  []interface{}
		expected bool
	}{
		{
			name:     "Example 1",
			root:     []interface{}{3, 4, 5, 1, 2},
			subRoot:  []interface{}{4, 1, 2},
			expected: true,
		},
		{
			name:     "Not a subtree",
			root:     []interface{}{1, 2, 3},
			subRoot:  []interface{}{2, 3},
			expected: false,
		},
		{
			name:     "Both empty",
			root:     []interface{}{},
			subRoot:  []interface{}{},
			expected: true,
		},
		{
			name:     "Subtree is empty",
			root:     []interface{}{1},
			subRoot:  []interface{}{},
			expected: true,
		},
		{
			name:     "Root is empty",
			root:     []interface{}{},
			subRoot:  []interface{}{1},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.root)
			sub := buildTree(tc.subRoot)
			got := isSubtree(root, sub)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
