package arrays

// combinationSum returns all unique combinations of candidates that sum to target.
func combinationSum(candidates []int, target int) [][]int {
	lst := [][]int{{}}
	for _, c := range candidates {
		for _, v := range lst {
			s := sum(v)
			tmp := make([]int, len(v))
			copy(tmp, v)
			for s <= target {
				tmp = append(tmp, c)
				lst = append(lst, tmp)
				s += c
			}
		}
	}

	result := [][]int{}
	for _, val := range lst {
		if sum(val) == target {
			result = append(result, val)
		}
	}

	return result
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
