package hackerrank

import "testing"

func TestDiagonalDifference(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name:     "Sample 1",
			input:    [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}},
			expected: 15,
		},
		{
			name:     "2x2 matrix",
			input:    [][]int{{1, 2}, {3, 4}},
			expected: 0,
		},
		{
			name:     "Negative values",
			input:    [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}},
			expected: 0,
		},
		{
			name:     "Single element",
			input:    [][]int{{5}},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DiagonalDifference(tc.input)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
