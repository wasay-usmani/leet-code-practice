package arrays

// characterReplacement returns the length of the longest substring with all the same letter after at most k replacements.
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	count := make([]int, 26)
	maxCount, left, result := 0, 0, 0
	for right := 0; right < len(s); right++ {
		idx := int(s[right] - 'A')
		count[idx]++
		if count[idx] > maxCount {
			maxCount = count[idx]
		}

		if right-left+1-maxCount > k {
			count[int(s[left]-'A')]--
			left++
		}
		
		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
