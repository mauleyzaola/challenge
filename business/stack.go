package business

import (
	"fmt"
	"strconv"
	"strings"
)

type numberStack struct {
	items []string
}

func (this *numberStack) Reverse() string {
	result := make([]string, len(this.items))
	for i := len(this.items) - 1; i >= 0; i-- {
		result = append(result, this.items[i])
	}
	return strings.Join(result, "")
}

func (this *numberStack) Value() string {
	return strings.Join(this.items, "")
}

func (this *numberStack) Float() float64 {
	value, _ := strconv.ParseFloat(this.Value(), 64)
	return value
}

func (this *numberStack) Push(value string) *numberStack {
	this.items = append(this.items, value)
	return this
}

func (this *numberStack) Pop() (string, error) {
	if len(this.items) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	defer func() {
		this.items = this.items[:len(this.items)-2]
	}()
	return this.items[len(this.items)-1], nil
}
