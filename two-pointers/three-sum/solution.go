package twopointers

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
