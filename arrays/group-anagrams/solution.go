package groupanagrams

import (
	"sort"
)

// GroupAnagrams groups anagrams together from the input slice.
func GroupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}
	for _, v := range strs {
		sorted := sortString(v)
		anagrams[sorted] = append(anagrams[sorted], v)
	}

	ret := [][]string{}
	for _, value := range anagrams {
		ret = append(ret, value)
	}

	return ret
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
