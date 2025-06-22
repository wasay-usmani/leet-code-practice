package trees

import (
	"reflect"
	"testing"
)

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

func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	var res []interface{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			res = append(res, nil)
			continue
		}
		res = append(res, n.Val)
		q = append(q, n.Left, n.Right)
	}
	// Trim trailing nils
	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}
	return res
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Example 1",
			input:    []interface{}{4, 2, 7, 1, 3, 6, 9},
			expected: []interface{}{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Empty",
			input:    []interface{}{},
			expected: nil,
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: []interface{}{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.input)
			inv := invertTree(root)
			got := treeToSlice(inv)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
