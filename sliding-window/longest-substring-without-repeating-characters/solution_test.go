package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{"Example 1", "abcabcbb", 3},
		{"Example 2", "bbbbb", 1},
		{"Example 3", "pwwkew", 3},
		{"Empty", "", 0},
		{"All unique", "abcdef", 6},
		{"Spaces", "a b c a b c", 3},
		{"Long repeat", "abba", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LengthOfLongestSubstring(tc.s)
			if result != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
