package binarysearch

// findMin returns the minimum element in a rotated sorted array.
func findMin(nums []int) int {
	var l, r, min, mid int = 0, len(nums) - 1, nums[0],0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < min {
			min = nums[mid]
		}

		if nums[l] > nums[r] && nums[l] < nums[mid] {
			l = mid + 1
		} else if nums[l] > nums[r] && nums[l] > nums[mid] {
			r = mid - 1
		} else if nums[l] < nums[r] && nums[r] > nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return min
}
