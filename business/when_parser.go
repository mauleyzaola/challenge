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
	return func(item *domain.BasketItem) (float64, error) {
		constants := make(map[string]float64)
		constants["price"] = item.Product.Price
		matches := float64(item.Quantity / number)
		noMatches := float64(item.Quantity) - matches

		eval, err := eval(priceExpr, constants)
		if err != nil {
			return 0, err
		}
		return (eval * matches) + (item.Product.Price * noMatches), nil
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
	return func(item *domain.BasketItem) (float64, error) {
		constants := make(map[string]float64)
		matches := false
		switch expr {
		case "gt":
			matches = item.Quantity > number
		case "gte":
			matches = item.Quantity >= number
		case "lt":
			matches = item.Quantity < number
		case "lte":
			matches = item.Quantity <= number
		}
		if matches {
			constants["price"] = item.Product.Price
			eval, err := eval(priceExpr, constants)
			if err != nil {
				return 0, err
			}
			return eval * float64(item.Quantity), nil
		} else {
			return item.Product.Price * float64(item.Quantity), nil
		}
	}, nil
}
