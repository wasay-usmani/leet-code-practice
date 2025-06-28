package hackerrank

import (
	"strconv"
)

// superDigit returns the super digit of the number created by concatenating n to itself k times.
func superDigit(n string, k int) int {
	super := supDigit(n)
	super = super * k
	return supDigit(strconv.Itoa(super))
}

func supDigit(n string) int {
	number, _ := strconv.ParseInt(n, 10, 64)
	for number > 9 {
		s := strconv.FormatInt(number, 10)
		sum := int64(0)
		for _, v := range s {
			sum = sum + int64(v-'0')
		}

		number = sum
	}

	return int(number)
}
