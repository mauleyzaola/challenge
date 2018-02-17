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

func isNumber(value uint8) bool {
	return value >= '0' && value <= '9'
}

// Many features are missing on this basic eval function, for instance operator precedence, grouping and making use of external functions is not implemented yet
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

func infixToPostfix(s string) (string, error) {
	var (
		curr   uint8
		number string
	)
	mainStack := newStringStack()
	auxStack := newStringStack()

	addAndClearNumber := func() {
		if len(number) != 0 {
			mainStack.Push(number)
			number = ""
		}
	}

	for i := range s {
		curr = s[i]
		if isNumber(curr) {
			number += string(curr)
		} else if isOperator(curr) {
			addAndClearNumber()
			if auxStack.Len() == 0 {
				auxStack.Push(string(curr))
			} else {
				if operatorOrder(string(curr)) >= operatorOrder(auxStack.Top()) {
					auxStack.Push(string(curr))
				} else {
					lastOp, _ := auxStack.Pop()
					mainStack.Push(lastOp)
				}
			}
		} else if isParenthesis(curr) {
			if curr == '(' {
				auxStack.Push(string(curr))
			} else {
				for auxStack.Len() != 0 {
					lastOp, err := auxStack.Pop()
					if err != nil {
						return "", err
					}
					if lastOp == ")" {
						continue
					} else if lastOp == "(" {
						continue
					} else {
						mainStack.Push(lastOp)
					}
				}
			}
		}
	}

	for auxStack.Len() != 0 {
		lastOp, _ := auxStack.Pop()
		mainStack.Push(lastOp)
	}

	// debug
	fmt.Println("input:", s)
	fmt.Println("main stack:", mainStack.Debug())
	fmt.Println("aux stack:", auxStack.Debug())

	return "", nil
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
