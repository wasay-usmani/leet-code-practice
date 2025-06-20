package twopointers

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{"Example 1", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{"Example 2", []int{0, 1, 1}, [][]int{}},
		{"Example 3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"No solution", []int{1, 2, 3}, [][]int{}},
		{"All zeros", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ThreeSum(tc.nums)
			if !tripletsEqual(result, tc.expected) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

// Helper to check if two slices of int triplets are equal regardless of order
func tripletsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, tripletA := range a {
		found := false
		for j, tripletB := range b {
			if !used[j] && reflect.DeepEqual(tripletA, tripletB) {
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
