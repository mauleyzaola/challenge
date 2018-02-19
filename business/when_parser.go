package business

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mauleyzaola/challenge/domain"
)

// Parses the WHEN expression and returns a callback function that can evaluate which indexes of the slice match the WHEN rule
//
// priceExpr is a formula that resolves the final price for each matched product
//
// It is up to the caller to decide what to do with the matched indexes
func WhenParser(whenExpr, priceExpr string) (domain.WhenCallback, error) {
	parseInt := func(value string) (int, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		return int(v), err
	}
	var (
		number int
		err    error
	)
	values := strings.Split(whenExpr, ":")
	switch values[0] {
	case "each":
		if len(values) != 2 {
			return nil, fmt.Errorf("expected 2 values for whenExpr but got instead:%d", len(values))
		}
		if number, err = parseInt(values[1]); err != nil {
			return nil, fmt.Errorf("cannot parse value:%s", values[1])
		}
		return whenEach(number, priceExpr), nil
	case "gte":
		fallthrough
	case "gt":
		if len(values) != 2 {
			return nil, fmt.Errorf("expected 2 values for whenExpr but got instead:%d", len(values))
		}
		if number, err = parseInt(values[1]); err != nil {
			return nil, fmt.Errorf("cannot parse value:%s", values[1])
		}
		return whenTotalCounter(number, values[0], priceExpr)
	}
	return nil, fmt.Errorf("cannot find any match for:%s", values[0])
}

func whenEach(number int, priceExpr string) domain.WhenCallback {
	return func(codes []string, products domain.Products) (float64, error) {
		var (
			err           error
			price, result float64
			items         domain.BasketItems
		)

		items, err = items.ToItems(codes, products)
		if err != nil {
			return 0, err
		}

		prMap, err := products.ToMap()
		if err != nil {
			return 0, err
		}
		counters := items.CountProducts()
		constants := make(map[string]float64)
		for code, count := range counters {
			price = prMap[code].Price
			constants["price"] = price
			matches := float64(count / number)
			noMatches := float64(count) - matches

			eval, err := Eval(priceExpr, constants)
			if err != nil {
				return 0, err
			}
			result += (eval * matches) + (price * noMatches)
		}
		return result, nil
	}
}

func whenTotalCounter(number int, expr, priceExpr string) (domain.WhenCallback, error) {
	switch expr {
	case "gt":
	case "gte":
	case "lt":
	case "lte":
	default:
		return nil, fmt.Errorf("unsupported expression:%s", expr)
	}
	return func(codes []string, products domain.Products) (float64, error) {
		var (
			err           error
			price, result float64
			items         domain.BasketItems
		)

		items, err = items.ToItems(codes, products)
		if err != nil {
			return 0, err
		}

		prMap, err := products.ToMap()
		if err != nil {
			return 0, err
		}

		counters := items.CountProducts()
		constants := make(map[string]float64)
		for code, count := range counters {
			matches := false
			switch expr {
			case "gt":
				matches = count > number
			case "gte":
				matches = count >= number
			case "lt":
				matches = count < number
			case "lte":
				matches = count <= number
			}
			price = prMap[code].Price
			if matches {
				constants["price"] = price
				eval, err := Eval(priceExpr, constants)
				if err != nil {
					return 0, err
				}
				result += eval * float64(count)
			} else {
				result += price * float64(count)
			}
		}

		return result, nil
	}, nil
}
