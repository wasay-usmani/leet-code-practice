package hackerrank

import "testing"

func TestSuperDigit(t *testing.T) {
	tests := []struct {
		name   string
		n      string
		k      int
		expect int
	}{
		{"example1", "148", 3, 3},
		{"single digit", "5", 1, 5},
		{"large k", "9875", 4, 8},
		{"all same digit", "1", 100000, 1},
		{"long n", "123", 3, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := superDigit(tt.n, tt.k)
			if got != tt.expect {
				t.Errorf("superDigit(%q, %d) = %d, want %d", tt.n, tt.k, got, tt.expect)
			}
		})
	}
}
