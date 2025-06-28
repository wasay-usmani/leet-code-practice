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

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		expect []int
	}{
		{"balanced BST", []interface{}{4, 2, 6, 1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"left skewed", []interface{}{3, 2, nil, 1}, []int{1, 2, 3}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, []int{1, 2, 3}},
		{"single node", []interface{}{42}, []int{42}},
		{"empty", []interface{}{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := InorderTraversal(root)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		expect []int
	}{
		{"balanced BST", []interface{}{4, 2, 6, 1, 3, 5, 7}, []int{4, 2, 1, 3, 6, 5, 7}},
		{"left skewed", []interface{}{3, 2, nil, 1}, []int{3, 2, 1}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, []int{1, 2, 3}},
		{"single node", []interface{}{42}, []int{42}},
		{"empty", []interface{}{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := PreorderTraversal(root)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("PreorderTraversal() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		expect []int
	}{
		{"balanced BST", []interface{}{4, 2, 6, 1, 3, 5, 7}, []int{1, 3, 2, 5, 7, 6, 4}},
		{"left skewed", []interface{}{3, 2, nil, 1}, []int{1, 2, 3}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, []int{3, 2, 1}},
		{"single node", []interface{}{42}, []int{42}},
		{"empty", []interface{}{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := PostorderTraversal(root)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		expect []int
	}{
		{"balanced BST", []interface{}{4, 2, 6, 1, 3, 5, 7}, []int{4, 2, 6, 1, 3, 5, 7}},
		{"left skewed", []interface{}{3, 2, nil, 1}, []int{3, 2, 1}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, []int{1, 2, 3}},
		{"single node", []interface{}{42}, []int{42}},
		{"empty", []interface{}{}, []int{}},
		{"zigzag", []interface{}{1, 2, 3, nil, 4, nil, 5}, []int{1, 2, 3, 4, 5}},
		{"unbalanced left", []interface{}{1, 2, nil, 3, nil, 4}, []int{1, 2, 3, 4}},
		{"unbalanced right", []interface{}{1, nil, 2, nil, 3, nil, 4}, []int{1, 2, 3, 4}},
		{"full but not complete", []interface{}{1, 2, 3, nil, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"root with only right child", []interface{}{1, nil, 2}, []int{1, 2}},
		{"root with only left child", []interface{}{1, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			got := LevelOrderTraversal(root)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("LevelOrderTraversal() = %v, want %v", got, tt.expect)
			}
		})
	}
}
