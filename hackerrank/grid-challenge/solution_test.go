package hackerrank

import "testing"

func TestGridChallenge(t *testing.T) {
	tests := []struct {
		name   string
		grid   []string
		expect string
	}{
		{"example1", []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}, "YES"},
		{"example2", []string{"abc", "lmp", "qrt"}, "YES"},
		{"example3", []string{"mpxz", "abcd", "wlmf"}, "NO"},
		{"single row", []string{"abc"}, "YES"},
		{"single col", []string{"a", "b", "c"}, "YES"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gridChallenge(tt.grid)
			if got != tt.expect {
				t.Errorf("gridChallenge(%v) = %q, want %q", tt.grid, got, tt.expect)
			}
		})
	}
}
