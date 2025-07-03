package applicationidparser

import (
	"strconv"
	"unicode"
)

// ParseApplicationIDs parses the input string and returns a slice of application IDs.
func ParseApplicationIDs(s string) []string {
	if s == "" {
		return []string{}
	}

	res := []string{}
	for i := 0; i < len(s)-1; i++ {
		prefix := ""
		for unicode.IsDigit(rune(s[i])) {
			prefix += string(s[i])
			i++
		}

		appLen, err := strconv.Atoi(prefix)
		if err != nil {
			panic(err)
		}

		res = append(res, s[i:i+appLen])
		i += appLen - 1
	}

	return res
}

// FilterWhitelisted filters the application IDs to only those in the whitelist.
func FilterWhitelisted(ids string, whitelist []string) []string {
	parsed := ParseApplicationIDs(ids)

	m := map[string]struct{}{}
	for _, v := range whitelist {
		m[v] = struct{}{}
	}

	res := []string{}
	for _, v := range parsed {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	return res
}
