package invertbinarytree

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
		if li < len(nodes) {
			node.Left = nodes[li]
		}
		if ri < len(nodes) {
			node.Right = nodes[ri]
		}
	}
	return nodes[0]
}

func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	result := []interface{}{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	// Trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected []interface{}
		name     string
	}{
		{[]interface{}{4, 2, 7, 1, 3, 6, 9}, []interface{}{4, 7, 2, 9, 6, 3, 1}, "example"},
		{[]interface{}{}, []interface{}{}, "empty"},
		{[]interface{}{2, 1, 3}, []interface{}{2, 3, 1}, "simple"},
		{[]interface{}{1, nil, 2}, []interface{}{1, 2}, "right child only"},
	}

	for _, tc := range tests {
		root := buildTree(tc.input)
		got := InvertTree(root)
		gotSlice := treeToSlice(got)
		if !reflect.DeepEqual(gotSlice, tc.expected) {
			t.Errorf("%s: got %v, want %v", tc.name, gotSlice, tc.expected)
		}
	}
}
