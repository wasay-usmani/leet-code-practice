package twopointers

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{
			name:     "Example 1",
			s1:       "abc",
			s2:       "ccccbbbbaaaa",
			expected: false,
		},
		{
			name:     "Example 2",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
		{
			name:     "Same string",
			s1:       "abc",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "Empty s1",
			s1:       "",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "Empty s2",
			s1:       "abc",
			s2:       "",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := checkInclusion(tc.s1, tc.s2)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
