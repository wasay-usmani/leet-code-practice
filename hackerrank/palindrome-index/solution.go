package hackerrank

// palindromeIndex returns the index of a character that can be removed to make s a palindrome, or -1 if not possible or already a palindrome.
func palindromeIndex(s string) int {
	isPal := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			if isPal(s, left+1, right) {
				return left
			}

			if isPal(s, left, right-1) {
				return right
			}

			return -1
		}

		left++
		right--
	}

	return -1
}
