package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "test 1",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "test 2",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "test 3",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Single element",
			nums:     []int{42},
			expected: false,
		},
		{
			name:     "All unique",
			nums:     []int{5, 6, 7, 8, 9},
			expected: false,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsDuplicate(tc.nums)
			if result != tc.expected {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
