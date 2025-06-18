package twosumii

import "testing"

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"Example 2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"Example 3", []int{-1, 0}, -1, []int{1, 2}},
		{"Large input", []int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSum(tc.numbers, tc.target)
			if len(result) != 2 || result[0] != tc.expected[0] || result[1] != tc.expected[1] {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
