package longestconsecutivesequence

import "slices"

// LongestConsecutive returns the length of the longest consecutive elements sequence.
func LongestConsecutive(nums []int) int {
	if len(nums) <= 0 {
		return len(nums)
	}

	slices.Sort(nums)
	var curr, max int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]-1 {
			curr += 1
		} else if nums[i] == nums[i+1] {
			continue
		} else {
			if curr > max {
				max = curr
			}

			curr = 0
		}
	}

	if curr > max {
		max = curr
	}

	return max + 1
}
