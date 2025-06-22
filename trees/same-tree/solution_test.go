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

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		p, q     []interface{}
		expected bool
	}{
		{
			name:     "Same trees",
			p:        []interface{}{1, 2, 3},
			q:        []interface{}{1, 2, 3},
			expected: true,
		},
		{
			name:     "Different structure",
			p:        []interface{}{1, 2},
			q:        []interface{}{1, nil, 2},
			expected: false,
		},
		{
			name:     "Different values",
			p:        []interface{}{1, 2, 1},
			q:        []interface{}{1, 1, 2},
			expected: false,
		},
		{
			name:     "Both empty",
			p:        []interface{}{},
			q:        []interface{}{},
			expected: true,
		},
		{
			name:     "One empty",
			p:        []interface{}{1},
			q:        []interface{}{},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pt := buildTree(tc.p)
			qt := buildTree(tc.q)
			got := isSameTree(pt, qt)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
