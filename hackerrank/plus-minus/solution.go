package hackerrank

import "fmt"

// PlusMinus prints the ratios of positive, negative, and zero values in the array to six decimal places.
func PlusMinus(arr []int) {
	var positive, negative, zero float32
	for _, v := range arr {
		switch {
		case v == 0:
			zero++
		case v < 0:
			negative++
		default:
			positive++
		}
	}

	l := float32(len(arr))
	fmt.Printf("%.6f\n", positive/l)
	fmt.Printf("%.6f\n", negative/l)
	fmt.Printf("%.6f\n", zero/l)
}
