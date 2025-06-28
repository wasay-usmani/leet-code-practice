package currencyconversion

import (
	"fmt"
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

var (
	currencyMap map[string][]Rate
)

// ParseCurrencies parses a currency conversion string into a map.
func ParseCurrencies(curr string) map[string][]Rate {
	currencyMap = make(map[string][]Rate)
	parsed := strings.Split(curr, comma)
	for _, v := range parsed {
		vals := strings.Split(v, colon)
		cost, _ := strconv.Atoi(vals[3])
		currencyMap[vals[0]] = append(currencyMap[vals[0]], Rate{To: vals[1], Delivery: vals[2], Cost: cost})
		currencyMap[vals[1]] = append(currencyMap[vals[1]], Rate{To: vals[0], Delivery: vals[2], Cost: cost})
	}

	return currencyMap
}

// ConvertCurrency returns the cost to convert from one currency to another.
func ConvertCurrency(from, to string) (int, error) {
	rates, ok := currencyMap[from]
	if !ok {
		return 0, fmt.Errorf("currency conversion from %s not found", from)
	}

	for _, rate := range rates {
		if rate.To == to {
			return rate.Cost, nil
		}
	}

	return 0, fmt.Errorf("currency conversion to %s not found", to)
}

type Res struct {
	Cost     int
	Delivery []string
}

// AtMostOneHop returns the cost and delivery path for at most one hop conversion.
func AtMostOneHop(from, to string) (*Res, error) {
	rates, ok := currencyMap[from]
	if !ok {
		return nil, fmt.Errorf("currency conversion from %s not found", from)
	}

	for _, rate := range rates {
		if rate.To == to {
			return &Res{Delivery: []string{rate.Delivery}, Cost: rate.Cost}, nil
		} else if rates2, ok := currencyMap[rate.To]; ok {
			for _, rate2 := range rates2 {
				if rate2.To == to {
					return &Res{Delivery: []string{rate.Delivery, rate2.Delivery}, Cost: rate.Cost + rate2.Cost}, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("currency conversion to %s not found", to)
}

// MinAtMostOneHop returns the minimum cost and delivery path for at most one hop conversion.
func MinAtMostOneHop(from, to string) (*Res, error) {
	rates, ok := currencyMap[from]
	if !ok {
		return nil, fmt.Errorf("currency conversion from %s not found", from)
	}

	var minCost int = 1000000000000
	var minRate *Res
	for _, rate := range rates {
		if rate.To == to {
			if rate.Cost < minCost {
				minCost = rate.Cost
				minRate = &Res{Delivery: []string{rate.Delivery}, Cost: rate.Cost}
			}
		} else if rates2, ok := currencyMap[rate.To]; ok {
			for _, rate2 := range rates2 {
				if rate2.To == to {
					if rate.Cost+rate2.Cost < minCost {
						minCost = rate.Cost + rate2.Cost
						minRate = &Res{Delivery: []string{rate.Delivery, rate2.Delivery}, Cost: rate.Cost + rate2.Cost}
					}
				}
			}
		}
	}

	return minRate, nil
}
