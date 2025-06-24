package currencyconversion

import (
	"errors"
	"strconv"
	"strings"
)

type Rate struct {
	To       string
	Delivery string
	Cost     int
}

const (
	comma = ","
	colon = ":"
)

func ParseCurrencies(curr string) (map[string][]Rate, error) {
	parts := strings.Split(curr, comma)
	m := make(map[string][]Rate)
	for _, v := range parts {
		s := strings.Split(v, colon)
		c, err := strconv.Atoi(s[3])
		if err != nil {
			return nil, errors.New("unable to parse")
		}

		m[s[0]] = append(m[s[0]], Rate{To: s[1], Delivery: s[2], Cost: c})
		m[s[1]] = append(m[s[1]], Rate{To: s[0], Delivery: s[2], Cost: c})
	}

	return m, nil
}

func ConvertCurrency(from, to string) (int, error) {
	fromMap, err := ParseCurrencies("USD:CAD:DHL:5,USD:GBP:FEDX:10")
	if err != nil {
		return 0, err
	}

	r, ok := fromMap[from]
	if !ok {
		return 0, errors.New("unable to convert")
	}

	for _, v := range r {
		if v.To == to {
			return v.Cost, nil
		}
	}

	return 0, errors.New("unable to convert")
}

type Res struct {
	Cost     int
	Delivery []string
}

func AtMostOneHop(from, to string) (*Res, error) {
	m, err := ParseCurrencies("USD:CAD:DHL:5,USD:GBP:FEDX:10")
	if err != nil {
		return nil, err
	}

	rates, ok := m[to]
	if !ok {
		return nil, errors.New("unable to convert")
	}

	for _, v := range rates {
		if v.To == from {
			return &Res{Cost: v.Cost, Delivery: []string{v.Delivery}}, nil
		} else {
			val := m[v.To]
			for _, rate := range val {
				if rate.To == from {
					return &Res{Cost: v.Cost + rate.Cost, Delivery: []string{v.Delivery, rate.Delivery}}, nil
				}
			}
		}
	}
	
	return nil, errors.New("unable to convert")
}

func MinAtMostOneHop(from, to string) (*Res, error) {
	m, err := ParseCurrencies("USD:CAD:DHL:5,USD:GBP:FEDX:10,GBP:CAD:FEDX:10")
	if err != nil {
		return nil, err
	}

	rates, ok := m[to]
	if !ok {
		return nil, errors.New("unable to convert")
	}

	result := []*Res{}
	for _, v := range rates {
		if v.To == from {
			result = append(result, &Res{Cost: v.Cost, Delivery: []string{v.Delivery}})
		}
		
		if v.To != from {
			val := m[v.To]
			for _, rate := range val {
				if rate.To == from {
					result = append(result, &Res{Cost: v.Cost + rate.Cost, Delivery: []string{v.Delivery, rate.Delivery}})
				}
			}
		}
	}

	min := 10000000000
	idx := 0
	for i, v := range result {
		if v.Cost < min {
			min = v.Cost
			idx = i
		}
	}

	return result[idx], nil
}
