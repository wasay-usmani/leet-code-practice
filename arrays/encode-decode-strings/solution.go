package encodedecodestrings

import (
	"strings"
)

var delim = ":;"

// Encode encodes a list of strings to a single string.
func Encode(strs []string) string {
	encoded := ""
	for _, v := range strs {
		encoded = encoded + v + delim
	}

	return encoded
}

// Decode decodes a single string to a list of strings.
func Decode(s string) []string {
	arr := strings.Split(s, delim)
	return arr[:len(arr)-1]
}
