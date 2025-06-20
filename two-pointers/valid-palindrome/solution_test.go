package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Example 1", "A man, a plan, a canal: Panama", true},
		{"Example 2", "race a car", false},
		{"Example 3", " ", true},
		{"Single char", "a", true},
		{"Numbers", "12321", true},
		{"Mixed", "No 'x' in Nixon", true},
		{"Not palindrome", "hello", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPalindrome(tc.s)
			if result != tc.expected {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
