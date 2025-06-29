package backtracking

import "sort"

// combinationSum2 returns all unique combinations in candidates where the numbers sum to target. Each number may only be used once.
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort the array to manage duplicates
	result := [][]int{}
	current := []int{}

	var backtrack func(target, start int)
	backtrack = func(target, start int) {
		if target == 0 {
			result = append(result, append([]int{}, current...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue // Skip duplicates
			}

			if candidates[i] > target {
				break // If the candidate exceeds the target, stop the loop
			}
			current = append(current, candidates[i])
			backtrack(target-candidates[i], i+1) // Move to the next candidate
			current = current[:len(current)-1]   // Backtrack by removing the last candidate
		}
	}

	backtrack(target, 0)
	return result
}
