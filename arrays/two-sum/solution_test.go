package arrays

import "testing"

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 3", []int{3, 3}, 6, []int{0, 1}},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"Large numbers", []int{1000000000, 2, 7, 1000000000}, 1000000002, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSum(tc.nums, tc.target)
			if !equalUnordered(result, tc.expected) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSumImproved(tc.nums, tc.target)
			if !equalUnordered(result, tc.expected) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

// Helper to check if two slices contain the same elements (order doesn't matter)
func equalUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
