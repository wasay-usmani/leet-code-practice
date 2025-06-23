package binarysearch

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "No rotation",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Single element",
			nums:     []int{10},
			expected: 10,
		},
		{
			name:     "Two elements",
			nums:     []int{2, 1},
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findMin(tc.nums)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
