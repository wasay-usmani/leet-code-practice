package groupanagrams

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Example 1",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "Example 2",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GroupAnagrams(tc.input)
			if !anagramGroupsEqual(result, tc.expected) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

// Helper to check if two slices of string slices are equal regardless of order
func anagramGroupsEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, groupA := range a {
		found := false
		for j, groupB := range b {
			if !used[j] && stringSliceEqualUnordered(groupA, groupB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func stringSliceEqualUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	ma := make(map[string]int)
	for _, s := range a {
		ma[s]++
	}
	for _, s := range b {
		if ma[s] == 0 {
			return false
		}
		ma[s]--
	}
	return true
}
