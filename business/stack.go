package business

import (
	"fmt"
)

type stringStack struct {
	_items []string
}

func newStringStack() *stringStack {
	return &stringStack{}
}

func (this *stringStack) len() int {
	return len(this._items)
}

func (this *stringStack) pop() (string, error) {
	if this.len() == 0 {
		return "", fmt.Errorf("empty stack")
	}
	defer func() {
		this._items = this._items[:len(this._items)-1]
	}()
	return this.top(), nil
}

func (this *stringStack) push(value string) {
	this._items = append(this._items, value)
}

func (this *stringStack) top() string {
	if this.len() == 0 {
		return ""
	}
	return this._items[len(this._items)-1]
}

func (this *stringStack) slice() []string {
	return this._items
}
