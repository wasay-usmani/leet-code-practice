package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{"Example 1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"Example 2", []int{4, 2, 0, 3, 2, 5}, 9},
		{"No water", []int{1, 2, 3, 4, 5}, 0},
		{"Flat", []int{0, 0, 0, 0}, 0},
		{"Single bar", []int{5}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Trap(tc.height)
			if result != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
