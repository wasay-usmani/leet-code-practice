package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
		name     string
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4, "example1"},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9, "example2"},
		{[]int{}, 0, "empty"},
		{[]int{1}, 1, "single"},
		{[]int{1, 2, 0, 1}, 3, "duplicates"},
		{[]int{10, 5, 12, 3, 55, 30, 4, 11, 2}, 4, "mixed"},
		{[]int{-2, -3, -1, 0, 1}, 5, "negatives"},
	}

	for _, tc := range tests {
		got := LongestConsecutive(tc.nums)
		if got != tc.expected {
			t.Errorf("%s: got %d, want %d", tc.name, got, tc.expected)
		}
	}
}
