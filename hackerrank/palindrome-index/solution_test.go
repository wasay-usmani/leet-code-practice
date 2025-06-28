package hackerrank

import "testing"

func TestPalindromeIndex(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		expect int
	}{
		{"remove last", "aaab", 3},
		{"remove first", "baa", 0},
		{"already palindrome", "aaa", -1},
		{"remove middle", "abcba", -1},
		{"remove to make palindrome", "babc", 3},
		{"no solution", "abc", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := palindromeIndex(tt.s)
			if got != tt.expect {
				t.Errorf("palindromeIndex(%q) = %d, want %d", tt.s, got, tt.expect)
			}
		})
	}
}
