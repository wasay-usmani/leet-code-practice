package hackerrank

import "testing"

func TestCountingSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sample 1",
			input:    []int{1, 1, 3, 2, 1},
			expected: func() []int { arr := make([]int, 100); arr[1] = 3; arr[2] = 1; arr[3] = 1; return arr }(),
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0},
			expected: func() []int { arr := make([]int, 100); arr[0] = 3; return arr }(),
		},
		{
			name:     "Max value",
			input:    []int{99, 99, 99},
			expected: func() []int { arr := make([]int, 100); arr[99] = 3; return arr }(),
		},
		{
			name:     "Empty",
			input:    []int{},
			expected: make([]int, 100),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountingSort(tc.input)
			if len(got) != 100 {
				t.Fatalf("expected length 100, got %d", len(got))
			}
			for i := 0; i < 100; i++ {
				if got[i] != tc.expected[i] {
					t.Errorf("at index %d: expected %d, got %d", i, tc.expected[i], got[i])
				}
			}
		})
	}
}
