package arrays

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			s:        "ABBB",
			k:        2,
			expected: 4,
		},
		{
			name:     "Example 2",
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			name:     "All same",
			s:        "AAAA",
			k:        2,
			expected: 4,
		},
		{
			name:     "No replacements",
			s:        "ABCD",
			k:        0,
			expected: 1,
		},
		{
			name:     "Empty string",
			s:        "",
			k:        2,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := characterReplacement(tc.s, tc.k)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
