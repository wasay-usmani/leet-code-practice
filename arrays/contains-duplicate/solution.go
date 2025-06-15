package containsduplicate

func ContainsDuplicate(nums []int) bool {
	visited := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := visited[v]; ok {
			return true
		}

		visited[v] = struct{}{}
	}

	return false
}
