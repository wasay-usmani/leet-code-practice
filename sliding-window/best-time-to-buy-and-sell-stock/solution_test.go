package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"Example 1", []int{7,1,5,3,6,4}, 5},
		{"Example 2", []int{7,6,4,3,1}, 0},
		{"Single day", []int{5}, 0},
		{"Increasing", []int{1,2,3,4,5}, 4},
		{"Decreasing", []int{5,4,3,2,1}, 0},
		{"Valley", []int{2,1,4}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MaxProfit(tc.prices)
			if result != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
