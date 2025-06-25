package arrays

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4, 1},
			expected: [][]int{{}, {3}, {2}, {2, 3}, {4}, {3, 4}, {2, 4}, {2, 3, 4}, {1}, {1, 3}, {1, 2}, {1, 2, 3}, {1, 4}, {1, 3, 4}, {1, 2, 4}, {1, 2, 3, 4}},
		},
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Empty",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: [][]int{{}, {5}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := subsets(tc.nums)
			if !equalUnordered(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

// Helper to compare two slices of slices regardless of order
func equalUnordered(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, x := range a {
		found := false
		for j, y := range b {
			if !used[j] && reflect.DeepEqual(x, y) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
