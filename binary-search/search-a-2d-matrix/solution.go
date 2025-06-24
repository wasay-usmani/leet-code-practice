package binarysearch

// searchMatrix returns true if target is found in the matrix.
func searchMatrix(matrix [][]int, target int) bool {
	var l, r int = 0, len(matrix) - 1
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
			return Search(matrix[mid], target)
		} else if target < matrix[mid][0] {
			r = mid - 1
		} else if target > matrix[mid][len(matrix[mid])-1] {
			l = mid + 1
		}
	}

	return false
}

func Search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
