package hackerrank

import "testing"

func TestLonelyInteger(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Sample 1",
			input:    []int{1, 2, 3, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "Single element",
			input:    []int{99},
			expected: 99,
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -2, -2, -1, -3},
			expected: -3,
		},
		{
			name:     "Large input",
			input:    []int{10, 20, 10, 30, 20, 40, 30},
			expected: 40,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LonelyInteger(tc.input)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
