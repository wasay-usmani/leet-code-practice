package attempt2

import (
	"strconv"
	"unicode"
)

func ParseApplicationIDs(s string) []string {
	if s == "" {
		return []string{}
	}

	ret := []string{}
	for i := 0; i < len(s)-1; i++ {
		prefix := ""
		for unicode.IsDigit(rune(s[i])) {
			prefix += string(s[i])
			i++
		}

		lenAppID, err := strconv.Atoi(prefix)
		if err != nil {
			panic(err)
		}

		ret = append(ret, s[i:i+lenAppID])
		i += lenAppID - 1
	}

	return ret
}

func FilterWhiteListed(s string, whitelisted []string) []string {
	if s == "" {
		return []string{}
	}

	m := map[string]struct{}{}
	for i := 0; i < len(s)-1; i++ {
		prefix := ""
		for unicode.IsDigit(rune(s[i])) {
			prefix += string(s[i])
			i++
		}

		lenAppID, err := strconv.Atoi(prefix)
		if err != nil {
			panic(err)
		}

		m[s[i:i+lenAppID]] = struct{}{}
		i += lenAppID - 1
	}

	ret := []string{}
	for _, v := range whitelisted {
		if _, ok := m[v]; ok {
			ret = append(ret, v)
		}
	}

	return ret
}
