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

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		expect bool
	}{
		{"example1", []interface{}{2, 1, 3}, true},
		{"example2", []interface{}{5, 1, 4, nil, nil, 3, 6}, false},
		{"single node", []interface{}{1}, true},
		{"empty", []interface{}{}, true},
		{"invalid left", []interface{}{10, 5, 15, nil, nil, 6, 20}, false},
		{"invalid right", []interface{}{10, 5, 15, nil, nil, 11, 9}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := isValidBST(root)
			if got != tt.expect {
				t.Errorf("isValidBST() = %v, want %v", got, tt.expect)
			}
		})
	}
}
