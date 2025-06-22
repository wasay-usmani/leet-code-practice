package hackerrank

// LonelyInteger returns the unique element in the array where all others occur twice.
func LonelyInteger(a []int) int {
	// Write your code here
	count := map[int]int{}
	for _, v := range a {
		count[v]++
	}

	for k, v := range count {
		if v == 1 {
			return k
		}
	}

	return 0
}
