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

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected int
	}{
		{
			name:     "Example 1",
			input:    []interface{}{3, 9, 20, nil, nil, 15, 7},
			expected: 3,
		},
		{
			name:     "Empty",
			input:    []interface{}{},
			expected: 0,
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: 1,
		},
		{
			name:     "Left skewed",
			input:    []interface{}{1, 2, nil, 3},
			expected: 3,
		},
		{
			name:     "Right skewed",
			input:    []interface{}{1, nil, 2, nil, nil, nil, 3},
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.input)
			got := maxDepth(root)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
