package arrays

import "testing"

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{"Example 1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"Example 2", []int{1}, 1, []int{1}},
		{"Negative numbers", []int{-1, -1, -2, -2, -2, -3}, 2, []int{-2, -1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TopKFrequent(tc.nums, tc.k)
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
