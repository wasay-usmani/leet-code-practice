package arrays

import "slices"

// subsets returns all possible subsets (the power set) of nums.
func subsets(nums []int) [][]int {
	set := [][]int{{}}
	for _, v := range nums {
		for _, setVal := range set {
			tmp := make([]int, len(setVal))
			copy(tmp, setVal)
            tmp = append(tmp, v)
			slices.Sort(tmp)
			set = append(set, tmp)
		}
	}

	return set
}