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

func TestPlusMinus(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []string
	}{
		{
			name:     "Sample 1",
			arr:      []int{-4, 3, -9, 0, 4, 1},
			expected: []string{"0.500000", "0.333333", "0.166667"},
		},
		{
			name:     "All positive",
			arr:      []int{1, 2, 3},
			expected: []string{"1.000000", "0.000000", "0.000000"},
		},
		{
			name:     "All negative",
			arr:      []int{-1, -2, -3},
			expected: []string{"0.000000", "1.000000", "0.000000"},
		},
		{
			name:     "All zero",
			arr:      []int{0, 0, 0},
			expected: []string{"0.000000", "0.000000", "1.000000"},
		},
		{
			name:     "Mixed",
			arr:      []int{1, -1, 0},
			expected: []string{"0.333333", "0.333333", "0.333333"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := captureOutput(func() { PlusMinus(tc.arr) })
			lines := strings.Split(strings.TrimSpace(output), "\n")
			if len(lines) != 3 {
				t.Fatalf("expected 3 lines, got %d", len(lines))
			}
			for i, exp := range tc.expected {
				if lines[i] != exp {
					t.Errorf("line %d: expected %s, got %s", i+1, exp, lines[i])
				}
			}
		})
	}
}
