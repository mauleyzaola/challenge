package business

import (
	"fmt"
	"strconv"
	"strings"
)

type stack struct {
	items []string
}

func (this *stack) reverse() string {
	result := make([]string, len(this.items))
	for i := len(this.items) - 1; i >= 0; i-- {
		result = append(result, this.items[i])
	}
	return strings.Join(result, "")
}

func (this *stack) value() string {
	return strings.Join(this.items, "")
}

func (this *stack) float() float64 {
	value, _ := strconv.ParseFloat(this.value(), 64)
	return value
}

func (this *stack) push(value string) *stack {
	this.items = append(this.items, value)
	return this
}

func (this *stack) pop() (string, error) {
	if len(this.items) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	defer func() {
		this.items = this.items[:len(this.items)-1]
	}()
	return this.items[len(this.items)-1], nil
}

func (this *stack) clear() {
	this.items = nil
}

func (this *stack) len() int {
	return len(this.items)
}
