package hackerrank

import "testing"

func TestFlippingMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int32
		expected int32
	}{
		{
			name:     "Sample 1",
			matrix:   [][]int32{{112, 42, 83, 119}, {56, 125, 56, 49}, {15, 78, 101, 43}, {62, 98, 114, 108}},
			expected: 414,
		},
		{
			name:     "2x2 matrix",
			matrix:   [][]int32{{1, 2}, {3, 4}},
			expected: 4,
		},
		{
			name:     "All same",
			matrix:   [][]int32{{5, 5, 5, 5}, {5, 5, 5, 5}, {5, 5, 5, 5}, {5, 5, 5, 5}},
			expected: 20,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FlippingMatrix(tc.matrix)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
