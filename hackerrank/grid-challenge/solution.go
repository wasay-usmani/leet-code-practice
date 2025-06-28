package hackerrank

import "slices"

// gridChallenge determines if the grid can be rearranged so that each row and each column is sorted alphabetically.
func gridChallenge(grid []string) string {
	sorted := make([]string, len(grid[0]))
	for _, s := range grid {
		r := []rune(s)
		slices.Sort(r)
		for j, b := range r {
			sorted[j] += string(b)
		}
	}

	for _, s := range sorted {
		r := []rune(s)
		if !slices.IsSorted(r) {
			return "NO"
		}
	}

	return "YES"
}
