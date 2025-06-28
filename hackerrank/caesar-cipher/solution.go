package hackerrank

import "unicode"

// caesarCipher encrypts the string s by shifting each letter by k positions.
func caesarCipher(s string, k int32) string {
	cipher := []rune{}
	for _, char := range s {
		if unicode.IsLetter(char) {
			if unicode.IsLower(char) {
				char = (char-'a'+k)%26 + 'a'
			} else {
				char = (char-'A'+k)%26 + 'A'
			}
		}

		cipher = append(cipher, char)
	}

	return string(cipher)
}
