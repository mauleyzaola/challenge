package business

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type stringStack struct {
	_stack *stack.Stack
}

func newStringStack() *stringStack {
	return &stringStack{_stack: &stack.Stack{}}
}

func (this *stringStack) Len() int {
	return this._stack.Len()
}

func (this *stringStack) Pop() (string, error) {
	if this.Len() == 0 {
		return "", fmt.Errorf("empty stack")
	}
	value, _ := this._stack.Pop().(string)
	return value, nil
}

func (this *stringStack) Push(value string) {
	this._stack.Push(value)
}

func (this *stringStack) Top() string {
	if value, ok := this._stack.Peek().(string); ok {
		return value
	}
	return ""
}
