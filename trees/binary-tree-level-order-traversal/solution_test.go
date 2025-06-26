package trees

import (
	"reflect"
	"testing"
)

func buildTreeLevelOrder(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	idx := 1
	for idx < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if idx < len(vals) && vals[idx] != nil {
			node.Left = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Left)
		}
		idx++
		if idx < len(vals) && vals[idx] != nil {
			node.Right = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		expect [][]int
	}{
		{
			"example1",
			[]interface{}{3, 9, 20, nil, nil, 15, 7},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			"single node",
			[]interface{}{1},
			[][]int{{1}},
		},
		{
			"empty",
			[]interface{}{},
			[][]int{},
		},
		{
			"left skewed",
			[]interface{}{1, 2, nil, 3},
			[][]int{{1}, {2}, {3}},
		},
		{
			"right skewed",
			[]interface{}{1, nil, 2, nil, 3},
			[][]int{{1}, {2}, {3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := levelOrder(root)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.expect)
			}
		})
	}
}
