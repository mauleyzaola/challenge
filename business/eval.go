package business

import (
	"fmt"
)

const (
	op_add = iota
	op_subs
	op_mul
	op_div
)

var ops map[string]int = map[string]int{
	"+": op_add,
	"-": op_subs,
	"*": op_mul,
	"/": op_div,
}

func isNumber(c uint8) bool {
	return c >= '0' && c <= '9'
}

// Many features are missing on this basic eval function, for instance operator precedence, grouping and making use of external functions is not implemented yet
func eval(expr string, constants map[string]float64) (float64, error) {
	return -1, fmt.Errorf("not implemented yet")
}

func isOperator(value uint8) bool {
	_, err := operator(value)
	return err != nil
}

func operator(value uint8) (int, error) {
	if val, ok := ops[string(value)]; ok {
		return val, nil
	}

	return -1, fmt.Errorf("unsupported operator:%s", value)
}

func toPostfix(s string) string {
	return ""
}

func operatorOrder(op string) int {
	switch op {
	case "+":
		fallthrough
	case "-":
		return 1
	case "*":
		fallthrough
	case "/":
		return 2
	default:
		return 0
	}
}
