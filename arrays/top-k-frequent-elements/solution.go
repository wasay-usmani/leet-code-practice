package topkfrequentelements

import (
	"sort"
)

// TopKFrequent returns the k most frequent elements in nums.
func TopKFrequent(nums []int, k int) []int {
	numberCount := make(map[int]int)
	for _, v := range nums {
		numberCount[v]++
	}

	tmp := []Entry{}
	for num, count := range numberCount {
		tmp = append(tmp, Entry{number: num, count: count})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].count > tmp[j].count // Descending by count
	})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, tmp[i].number)
	}

	return result
}

type Entry struct {
	number, count int
}
