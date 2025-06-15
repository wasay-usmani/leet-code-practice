package twosum

// TwoSum returns indices of the two numbers such that they add up to target.
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func TwoSumImproved(nums []int, target int) []int {
	differences := map[int]int{}
	for i, v := range nums {
		if index, ok := differences[target-v]; ok {
			return []int{i, index}
		}

		differences[v] = i
	}

	return []int{}
}
