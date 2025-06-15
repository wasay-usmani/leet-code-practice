package validanagram

// IsAnagram returns true if t is an anagram of s.
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	words := map[rune]int{}
	for _, key := range s {
		words[key]++
	}

	for _, key := range t {
		if val, ok := words[key]; !ok || val == 0 {
			return false
		}

		words[key] = words[key] - 1
	}

	return true
}
