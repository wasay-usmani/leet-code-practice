package attempt3

import (
	"strconv"
	"unicode"
)

func ParseApplicationIDs(s string) []string {
	res := []string{}
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

		res = append(res, s[i:i+lenAppID])
		i += lenAppID - 1
	}

	return res
}

func ParseWhiteListed(s string, whitelisted []string) []string {
	idMap := map[string]struct{}{}
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

		idMap[s[i:i+lenAppID]] = struct{}{}
		i += lenAppID - 1
	}

	res := []string{}
	for _, v := range whitelisted {
		if _, ok := idMap[v]; ok {
			res = append(res, v)
		}
	}

	return res
}
