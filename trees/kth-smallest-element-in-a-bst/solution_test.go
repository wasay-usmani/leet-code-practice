package trees

import (
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

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		k      int
		expect int
	}{
		{"example1", []interface{}{3, 1, 4, nil, 2}, 1, 1},
		{"example2", []interface{}{3, 1, 4, nil, 2}, 2, 2},
		{"example3", []interface{}{5, 3, 6, 2, 4, nil, nil, 1}, 3, 3},
		{"single node", []interface{}{1}, 1, 1},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, 2, 2},
		{"left skewed", []interface{}{3, 2, nil, 1}, 2, 2},
		{"k is last", []interface{}{2, 1, 3}, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := kthSmallest(root, tt.k)
			if got != tt.expect {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.expect)
			}
		})
	}
}
