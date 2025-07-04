package attempt3

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Conversion struct {
	To       string
	Delivery string
	Cost     int64
}

const (
	comma = ","
	colon = ":"
)

func ParseConversions(s string) map[string][]Conversion {
	conversions := strings.Split(s, comma)
	convMap := map[string][]Conversion{}
	for _, conv := range conversions {
		val := strings.Split(conv, colon)
		c, err := strconv.ParseInt(val[3], 10, 64)
		if err != nil {
			panic(err)
		}

		convMap[val[0]] = append(convMap[val[0]], Conversion{To: val[1], Delivery: val[2], Cost: c})
		convMap[val[1]] = append(convMap[val[1]], Conversion{To: val[0], Delivery: val[2], Cost: c})
	}

	return convMap
}

func Convert(s, to, from string) (int64, error) {
	convMap := ParseConversions(s)
	curr, ok := convMap[from]
	if !ok {
		return 0, errors.New("unable to parse conversions")
	}

	for _, c := range curr {
		if c.To == to {
			return c.Cost, nil
		}
	}

	return 0, errors.New("currency conversion not supported")
}

type Res struct {
	Delivery []string
	Cost     int64
}

func AtMostOneHop(s, to, from string) (*Res, error) {
	convMap := ParseConversions(s)
	curr, ok := convMap[from]
	if !ok {
		return nil, errors.New("unable to parse conversions")
	}

	for _, c := range curr {
		if c.To == to {
			return &Res{Cost: c.Cost, Delivery: []string{c.Delivery}}, nil
		} else {
			oneHop := convMap[c.To]
			for _, v := range oneHop {
				if v.To == to {
					return &Res{Cost: c.Cost + v.Cost, Delivery: []string{c.Delivery, v.Delivery}}, nil
				}
			}
		}
	}

	return nil, errors.New("currency conversion not supported")
}

func MinCost(s, to, from string) (*Res, error) {
	convMap := ParseConversions(s)
	curr, ok := convMap[from]
	if !ok {
		return nil, errors.New("unable to parse conversions")
	}

	var min int64 = math.MaxInt64
	var res *Res
	for _, c := range curr {
		if c.To == to {
			if c.Cost < min {
				res = &Res{Cost: c.Cost, Delivery: []string{c.Delivery}}
			}
		} else {
			oneHop := convMap[c.To]
			for _, v := range oneHop {
				if v.To == to {
					if c.Cost < min {
						res = &Res{Cost: c.Cost + v.Cost, Delivery: []string{c.Delivery, v.Delivery}}
					}
				}
			}
		}
	}

	return res, nil
}
