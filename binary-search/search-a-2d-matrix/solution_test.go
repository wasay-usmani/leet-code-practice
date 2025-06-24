package binarysearch

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name:     "Example 1",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			name:     "Example 2",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
		{
			name:     "Single row",
			matrix:   [][]int{{1, 2, 3, 4}},
			target:   3,
			expected: true,
		},
		{
			name:     "Single column",
			matrix:   [][]int{{1}, {3}, {5}},
			target:   4,
			expected: false,
		},
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			target:   1,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := searchMatrix(tc.matrix, tc.target)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
