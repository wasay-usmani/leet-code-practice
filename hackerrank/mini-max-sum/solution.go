package hackerrank

import "fmt"

// MiniMaxSum prints the minimum and maximum sum of 4 out of 5 integers in arr.
func MiniMaxSum(arr []int) {
	var min, max, total int
	for _, v := range arr {
		total += v
	}

	min = total
	for _, v := range arr {
		if total-v > max {
			max = total - v
		}

		if total-v < min {
			min = total - v
		}
	}

	fmt.Printf("%d %d", min, max)
}
