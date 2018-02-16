package business

import (
	"fmt"
	"strconv"
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

// Many features are missing on this basic eval function, for instance operator precedence, grouping and making use of external functions is not implemented yet
func evaluate(expr string, constants map[string]float64) (float64, error) {
	var result float64
	constStack, opStack, numberStack := &stack{}, &stack{}, &stack{}

	for _, v := range expr {
		value := string(v)
		if isOperator(value) {
			if opStack.len() != 0 {
				return -1, fmt.Errorf("two continuous operators are not supported")
			}
			opStack.push(value)
		} else if isNumber(value) {
			numberStack.push(value)
		} else { // assume it is part of a constant
			constStack.push(value)
		}
	}

	return result, nil
}

func isNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func isOperator(value string) bool {
	_, err := operator(value)
	return err != nil
}

func operator(value string) (int, error) {
	if val, ok := ops[value]; ok {
		return val, nil
	}

	return -1, fmt.Errorf("unsupported operator:%s", value)
}
