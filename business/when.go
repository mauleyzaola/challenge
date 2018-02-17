package business

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mauleyzaola/challenge/domain"
)

// Evaluates the WHEN expression and returns a callback function that can eval which indexes of the slice match the WHEN rule
//
// It is up to the caller to decide what to do with the matched indexes
func WhenFilter(when string) (domain.WhenCallback, error) {
	parseInt := func(value string) (int, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		return int(v), err
	}
	var (
		number int
		err    error
	)
	values := strings.Split(when, ":")
	switch values[0] {
	case "each":
		if len(values) != 2 {
			return nil, fmt.Errorf("expected 2 values for when but got instead:%d", len(values))
		}
		if number, err = parseInt(values[1]); err != nil {
			return nil, fmt.Errorf("cannot parse value:%s", values[1])
		}
		return whenEach(number), nil
	case "gte":
		fallthrough
	case "gt":
		if len(values) != 2 {
			return nil, fmt.Errorf("expected 2 values for when but got instead:%d", len(values))
		}
		if number, err = parseInt(values[1]); err != nil {
			return nil, fmt.Errorf("cannot parse value:%s", values[1])
		}
		return whenTotalCounter(number, values[0])
	}
	return nil, fmt.Errorf("cannot find any match for:%s", values[0])
}

func whenEach(number int) domain.WhenCallback {
	return func(products []domain.Product) map[int]bool {
		result := make(map[int]bool)
		for i := range products {
			if i == 0 {
				continue
			}
			if i%number == 0 {
				result[i] = true
			}
		}
		return result
	}
}

func whenTotalCounter(number int, expr string) (domain.WhenCallback, error) {
	switch expr {
	case "gt":
	case "gte":
	case "lt":
	case "lte":
	default:
		return nil, fmt.Errorf("unsupported expression:%s", expr)
	}
	return func(products []domain.Product) map[int]bool {
		result := make(map[int]bool)
		matches := false
		switch expr {
		case "gt":
			matches = len(products) > number
		case "gte":
			matches = len(products) >= number
		case "lt":
			matches = len(products) < number
		case "lte":
			matches = len(products) <= number
		}
		if matches {
			for i := range products {
				result[i] = true
			}
		}
		return result
	}, nil
}
