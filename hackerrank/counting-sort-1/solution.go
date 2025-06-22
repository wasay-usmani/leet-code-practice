package hackerrank

// CountingSort returns a frequency array for the values in arr (range 0-99).
func CountingSort(arr []int) []int {
	result := make([]int, 100)
	for _, v := range arr {
		result[v]++
	}

	return result
}
