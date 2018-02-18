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

func isNumber(value uint8) bool {
	return (value >= '0' && value <= '9') || value == '.'
}

func eval(expr string, constants map[string]float64) (float64, error) {
	return -1, fmt.Errorf("not implemented yet")
}

func isOperator(value uint8) bool {
	_, err := operator(value)
	return err == nil
}

func isParenthesis(value uint8) bool {
	return value == '(' || value == ')'
}

func operator(value uint8) (int, error) {
	if val, ok := ops[string(value)]; ok {
		return val, nil
	}

	return -1, fmt.Errorf("unsupported operator:%s", value)
}

// TODO implement constants as well
func infixToPostfix(s string) (string, error) {
	var (
		curr uint8
	)
	mainStack := newStringStack()
	auxStack := newStringStack()

	for i := 0; i < len(s); i++ {
		curr = s[i]

		if isParenthesis(curr) {
			if curr == '(' {
				auxStack.Push(string(curr))
			} else { // closing parenthesis
				for auxStack.Len() != 0 {
					oldVal, _ := auxStack.Pop()
					if oldVal == "(" {
						continue
					}
					if oldVal == ")" {
						break
					}
					mainStack.Push(oldVal)
				}
			}
		} else if isOperator(curr) {
			if auxStack.Len() != 0 {
				if operatorOrder(string(curr)) > operatorOrder(auxStack.Top()) {
					auxStack.Push(string(curr))
				} else if operatorOrder(string(curr)) == operatorOrder(auxStack.Top()) {
					oldVal, _ := auxStack.Pop()
					mainStack.Push(oldVal)
					auxStack.Push(string(curr))
				}
			} else {
				auxStack.Push(string(curr))
			}
		} else if isNumber(curr) {
			// we need to consider numbers with more than one digit
			number := ""
			for j := i; j < len(s) && isNumber(s[j]); j++ {
				number += string(s[j])
			}
			// validate if the number can be parsed as float (cases like invalid decimal points for instance)
			if _, err := strconv.ParseFloat(number, 64); err != nil {
				return "", fmt.Errorf("invalid number:%s. %s", number, err)
			}
			mainStack.Push(number)
			i += len(number) - 1
		} else {
			return "", fmt.Errorf("not a valid number, operator or parenthesis:%s", string(curr))
		}
	}

	for auxStack.Len() != 0 {
		lastOp, _ := auxStack.Pop()
		mainStack.Push(lastOp)
	}

	return mainStack.Split(), nil
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
