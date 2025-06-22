package hackerrank

// DiagonalDifference returns the absolute difference between the sums of the matrix's two diagonals.
func DiagonalDifference(arr [][]int) int {
	var left, right int
	for i := 0; i < len(arr); i++ {
		left += arr[i][i]
		right += arr[i][len(arr)-i-1]
	}

	return AbsInt(right - left)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
