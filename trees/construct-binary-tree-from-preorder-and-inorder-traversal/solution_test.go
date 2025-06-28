package main

import (
	"reflect"
	"testing"
)

// Helper function to convert tree to slice representation (level order)
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var result []interface{}
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

	// Remove trailing nil values
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// Helper function to create a tree from slice representation
func sliceToTree(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	if values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []interface{}
	}{
		{
			name:     "Example 1: Basic tree",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:     "Example 2: Single node",
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: []interface{}{-1},
		},
		{
			name:     "Empty arrays",
			preorder: []int{},
			inorder:  []int{},
			expected: []interface{}{},
		},
		{
			name:     "Left skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			expected: []interface{}{1, 2, nil, 3},
		},
		{
			name:     "Balanced tree with 3 nodes",
			preorder: []int{2, 1, 3},
			inorder:  []int{1, 2, 3},
			expected: []interface{}{2, 1, 3},
		},
		{
			name:     "Complex tree",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildTree(tt.preorder, tt.inorder)
			resultSlice := treeToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("buildTree() = %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}

// Benchmark test
func BenchmarkBuildTree(b *testing.B) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	inorder := []int{4, 2, 5, 1, 6, 3, 7}

	for i := 0; i < b.N; i++ {
		buildTree(preorder, inorder)
	}
}

// Test edge cases
func TestBuildTreeEdgeCases(t *testing.T) {
	t.Run("Nil input handling", func(t *testing.T) {
		result := buildTree(nil, nil)
		if result != nil {
			t.Errorf("Expected nil for nil inputs, got %v", result)
		}
	})

	t.Run("Mismatched lengths", func(t *testing.T) {
		// This should handle gracefully or panic appropriately
		defer func() {
			if r := recover(); r != nil {
				// Expected behavior for invalid input
			}
		}()

		buildTree([]int{1, 2}, []int{1})
	})

	t.Run("Single element arrays", func(t *testing.T) {
		result := buildTree([]int{42}, []int{42})
		expected := []interface{}{42}

		if !reflect.DeepEqual(treeToSlice(result), expected) {
			t.Errorf("Expected %v, got %v", expected, treeToSlice(result))
		}
	})
}
