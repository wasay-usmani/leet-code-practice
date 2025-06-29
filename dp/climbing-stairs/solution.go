package dp

// climbStairs returns the number of distinct ways to climb to the top.
func climbStairs(n int) int {
	one, two := 1, 1
	for range n - 1 {
		one, two = two, one+two
	}

	return two
}
