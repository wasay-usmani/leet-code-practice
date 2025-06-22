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

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected bool
	}{
		{
			name:     "Right heavy",
			input:    []interface{}{1, nil, 2, nil, nil, nil, 3},
			expected: false,
		},
		{
			name:     "Example 1",
			input:    []interface{}{3, 9, 20, nil, nil, 15, 7},
			expected: true,
		},
		{
			name:     "Unbalanced",
			input:    []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4},
			expected: false,
		},
		{
			name:     "Empty",
			input:    []interface{}{},
			expected: true,
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: true,
		},
		{
			name:     "Left heavy",
			input:    []interface{}{1, 2, nil, 3},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.input)
			got := isBalanced(root)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
