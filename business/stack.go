package business

import (
	"fmt"
	"strings"
)

type stringStack struct {
	_items []string
}

func newStringStack() *stringStack {
	return &stringStack{}
}

func (this *stringStack) Len() int {
	return len(this._items)
}

func (this *stringStack) Pop() (string, error) {
	if this.Len() == 0 {
		return "", fmt.Errorf("empty stack")
	}
	defer func() {
		this._items = this._items[:len(this._items)-1]
	}()
	return this.Top(), nil
}

func (this *stringStack) Push(value string) {
	this._items = append(this._items, value)
}

func (this *stringStack) Top() string {
	if this.Len() == 0 {
		return ""
	}
	return this._items[len(this._items)-1]
}

func (this *stringStack) Debug() string {
	return strings.Join(this._items, ",")
}
