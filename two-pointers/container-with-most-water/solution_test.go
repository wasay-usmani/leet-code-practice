package twopointers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
		name     string
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49, "example1"},
		{[]int{1, 1}, 1, "example2"},
		{[]int{4, 3, 2, 1, 4}, 16, "symmetric"},
		{[]int{1, 2, 1}, 2, "valley"},
		{[]int{2, 3, 10, 5, 7, 8, 9}, 36, "increasing"},
		{[]int{1, 2, 4, 3}, 4, "small"},
	}

	for _, tc := range tests {
		got := MaxArea(tc.height)
		if got != tc.expected {
			t.Errorf("%s: got %d, want %d", tc.name, got, tc.expected)
		}
	}
}
