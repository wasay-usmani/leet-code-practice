package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"Example 1", "anagram", "nagaram", true},
		{"Example 2", "rat", "car", false},
		{"Empty strings", "", "", true},
		{"Single char true", "a", "a", true},
		{"Single char false", "a", "b", false},
		{"Different lengths", "abc", "ab", false},
		{"Unicode false", "你好", "好你啊", false},
		{"Unicode true", "你好", "好你", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsAnagram(tc.s, tc.t)
			if result != tc.expected {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
