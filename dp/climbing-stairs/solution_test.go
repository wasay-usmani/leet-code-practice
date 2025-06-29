package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		expect int
	}{
		{"n=2", 2, 2},
		{"n=3", 3, 3},
		{"n=4", 4, 5},
		{"n=5", 5, 8},
		{"n=1", 1, 1},
		{"n=0", 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.expect {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.expect)
			}
		})
	}
}
