package hackerrank

import (
	"fmt"
	"strconv"
	"strings"
)

// TimeConversion converts a time in 12-hour AM/PM format to 24-hour format.
func TimeConversion(s string) string {
	parsed := strings.Split(s, ":")
	h, m, sec, pm := parsed[0], parsed[1], parsed[2][:2], string(parsed[2][2:4])
	switch pm {
	case "PM":
		if h != "12" {
			hour, _ := strconv.Atoi(h)
			h = strconv.Itoa(hour + 12)
		}
	case "AM":
		if h == "12" {
			h = "00"
		}
	}

	return fmt.Sprintf("%v:%v:%v", h, m, sec)
}
