package hackerrank

// FlippingMatrix returns the maximum sum of the n x n upper-left submatrix after flips.
func FlippingMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2
	var sum int32
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := matrix[i][j]
			b := matrix[i][len(matrix)-1-j]
			c := matrix[len(matrix)-1-i][j]
			d := matrix[len(matrix)-1-i][len(matrix)-1-j]

			sum += max(a, b, c, d)
		}
	}

	return sum
}
