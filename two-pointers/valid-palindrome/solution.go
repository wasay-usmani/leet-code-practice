package twopointers

import (
	"unicode"
)

// IsPalindrome returns true if s is a valid palindrome (ignoring non-alphanumeric and case).
func IsPalindrome(s string) bool {
	san := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			san = append(san, unicode.ToLower(v))
		}
	}

	for i := 0; i < len(san); i++ {
		if san[i] != san[len(san)-i-1] {
			return false
		}
	}

	return true
}
