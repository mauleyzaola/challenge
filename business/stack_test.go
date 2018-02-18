package business

import "testing"

func TestStringStack_Len(t *testing.T) {
	stack := newStringStack()
	stack.push("1")
	stack.push("20")
	stack.push("300")
	expected := 3
	result := stack.len()
	if expected != result {
		t.Errorf("expected:%d but got instead:%d", expected, result)
	}
}

func TestStringStack_Pop(t *testing.T) {
	stack := newStringStack()
	stack.push("1")
	stack.push("20")
	expected := "20"
	result, err := stack.pop()
	if err != nil {
		t.Error(err)
	}
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}

	expected = "1"
	result, err = stack.pop()
	if err != nil {
		t.Error(err)
	}
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}

	expected = ""
	result, err = stack.pop()
	if err == nil {
		t.Error("expected error but got nil instead")
	}
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}

func TestStringStack_Top(t *testing.T) {
	stack := newStringStack()
	stack.push("1")
	stack.push("20")

	expected := "20"
	result := stack.top()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}

	stack.pop()
	expected = "1"
	result = stack.top()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}
