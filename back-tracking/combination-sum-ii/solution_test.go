package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func sort2DIntSlice(s [][]int) {
	for _, v := range s {
		sort.Ints(v)
	}
	sort.Slice(s, func(i, j int) bool {
		for x := range s[i] {
			if x >= len(s[j]) {
				return false
			}
			if s[i][x] != s[j][x] {
				return s[i][x] < s[j][x]
			}
		}
		return len(s[i]) < len(s[j])
	})
}

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expect     [][]int
	}{
		{"example1", []int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"example2", []int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		{"no solution", []int{2, 4, 6}, 5, [][]int{}},
		{"single element", []int{3}, 3, [][]int{{3}}},
		{"empty", []int{}, 3, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			sort2DIntSlice(got)
			sort2DIntSlice(tt.expect)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.expect)
			}
		})
	}
}
