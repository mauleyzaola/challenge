package business

import "testing"

func TestStringStack_Len(t *testing.T) {
	stack := newStringStack()
	stack.Push("1")
	stack.Push("20")
	stack.Push("300")
	expected := 3
	result := stack.Len()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}

func TestStringStack_Pop(t *testing.T) {
	stack := newStringStack()
	stack.Push("1")
	stack.Push("20")
	expected := "20"
	result, err := stack.Pop()
	if err != nil {
		t.Error(err)
	}
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}

	expected = "1"
	result, err = stack.Pop()
	if err != nil {
		t.Error(err)
	}
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}

	expected = ""
	result, err = stack.Pop()
	if err == nil {
		t.Error("expected error but got nil instead")
	}
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}

func TestStringStack_Top(t *testing.T) {
	stack := newStringStack()
	stack.Push("1")
	stack.Push("20")

	expected := "20"
	result := stack.Top()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}

	stack.Pop()
	expected = "1"
	result = stack.Top()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}
