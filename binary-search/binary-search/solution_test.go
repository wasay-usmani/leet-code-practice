package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"Example 1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"Example 2", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"Single element found", []int{5}, 5, 0},
		{"Single element not found", []int{5}, 1, -1},
		{"First element", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Last element", []int{1, 2, 3, 4, 5}, 5, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Search(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
