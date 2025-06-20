package slidingwindow

// LengthOfLongestSubstring returns the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	visited := map[byte]struct{}{}
	l, r, m, count := 0, 0, 0, 0
	for r < len(s) {
		if _, ok := visited[s[r]]; ok {
			if count > m {
				m = count
			}

			for l < r {
				delete(visited, s[l])
				l++
				count--

				if s[l-1] == s[r] {
					break
				}
			}

		} else {

		}

		count++
		visited[s[r]] = struct{}{}
		r++
	}

	return max(count, m)
}
