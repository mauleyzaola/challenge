package business

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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
	_, err := parseOperator(value)
	return err == nil
}

func isParenthesis(value uint8) bool {
	return value == '(' || value == ')'
}

func parseOperator(value uint8) (int, error) {
	if val, ok := ops[string(value)]; ok {
		return val, nil
	}

	return -1, fmt.Errorf("unsupported parseOperator:%s", string(value))
}

func sortBySize(constants StringConstants) StringConstants {
	sort.Sort(constants)
	return constants
}

func sortConstants(constants map[string]float64) map[string]float64 {
	var slice []string
	for k := range constants {
		slice = append(slice, k)
	}
	slice = sortBySize(slice)
	result := make(map[string]float64, len(slice))
	for _, v := range slice {
		result[v] = constants[v]
	}
	return result
}

func infixToPostfix(expr string, constants map[string]float64) ([]string, error) {
	var curr uint8
	mainStack := newStringStack()
	auxStack := newStringStack()

	for i := 0; i < len(expr); i++ {
		curr = expr[i]

		if isParenthesis(curr) {
			if curr == '(' {
				auxStack.push(string(curr))
			} else { // closing parenthesis
				for auxStack.len() != 0 {
					last, _ := auxStack.pop()
					if last == "(" {
						continue
					}
					if last == ")" {
						break
					}
					mainStack.push(last)
				}
			}
		} else if isOperator(curr) {
			if auxStack.len() != 0 {
				if operatorOrder(string(curr)) > operatorOrder(auxStack.top()) {
					auxStack.push(string(curr))
				} else {
					last, _ := auxStack.pop()
					mainStack.push(last)
					auxStack.push(string(curr))
				}
			} else {
				auxStack.push(string(curr))
			}
		} else if isNumber(curr) {
			// we need to consider numbers with more than one digit
			number := ""
			for j := i; j < len(expr) && isNumber(expr[j]); j++ {
				number += string(expr[j])
			}
			// validate if the number can be parsed as float (cases like invalid decimal points for instance)
			if _, err := strconv.ParseFloat(number, 64); err != nil {
				return nil, fmt.Errorf("invalid number:%value. %value", number, err)
			}
			mainStack.push(number)
			i += len(number) - 1
		} else {
			return nil, fmt.Errorf("not a valid number, parseOperator or parenthesis:%value", string(curr))
		}
	}

	for auxStack.len() != 0 {
		last, _ := auxStack.pop()
		mainStack.push(last)
	}

	return mainStack.slice(), nil
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

func postfixCalculator(values []string) (float64, error) {
	stack := newStringStack()

	doCalc := func(op string) error {
		var result float64
		val, err := stack.pop()
		if err != nil {
			return err
		}
		number2, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		val, err = stack.pop()
		if err != nil {
			return err
		}
		number1, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		o, err := parseOperator(op[0])
		if err != nil {
			return err
		}
		switch o {
		case op_add:
			result = number1 + number2
		case op_subs:
			result = number1 - number2
		case op_mul:
			result = number1 * number2
		case op_div:
			if number2 == 0 {
				return fmt.Errorf("division by zero not supported")
			}
			result = number1 / number2
		}
		stack.push(formatFloat(result))
		return nil
	}

	// at this point we assume there are only numbers and operators
	for _, v := range values {
		if !isOperator(v[0]) {
			stack.push(v)
			continue
		}
		if err := doCalc(v); err != nil {
			return 0, err
		}
	}

	if stack.len() != 1 {
		return 0, fmt.Errorf("invalid postfix expression:%s", strings.Join(values, ""))
	}

	return strconv.ParseFloat(stack.slice()[0], 64)
}

func formatFloat(number float64) string {
	return strconv.FormatFloat(number, 'f', 6, 64)
}
