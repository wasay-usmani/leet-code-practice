package hackerrank

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestMiniMaxSum(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected string
	}{
		{
			name:     "Sample 1",
			arr:      []int{1, 2, 3, 4, 5},
			expected: "10 14",
		},
		{
			name:     "All same",
			arr:      []int{7, 7, 7, 7, 7},
			expected: "28 28",
		},
		{
			name:     "Descending order",
			arr:      []int{9, 8, 7, 6, 5},
			expected: "26 30",
		},
		{
			name:     "With zero",
			arr:      []int{0, 1, 2, 3, 4},
			expected: "6 10",
		},
		{
			name:     "Large numbers",
			arr:      []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000},
			expected: "4000000000 4000000000",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := strings.TrimSpace(captureOutput(func() { MiniMaxSum(tc.arr) }))
			if output != tc.expected {
				t.Errorf("expected '%s', got '%s'", tc.expected, output)
			}
		})
	}
}
