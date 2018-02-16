package business

import "testing"

func TestStackCast(t *testing.T) {
	stack := &stack{}
	stack.push("1").push("2")
	expected := "12"
	result := stack.value()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
	var fexpected float64 = 12
	fresult := stack.float()
	if fexpected != fresult {
		t.Errorf("expected:%v but got instead:%v", fexpected, fresult)
	}
}

func TestStackReverse(t *testing.T) {
	stack := &stack{}
	stack.push("123").push("456")
	expected := "456123"
	result := stack.reverse()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}

func TestStackSlice(t *testing.T) {
	stack := &stack{}
	stack.push("1").push("2")
	expected := "2"
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
	_, err = stack.pop()
	if err == nil {
		t.Error("expected error but got nil instead")
	}

	exp := 0
	res := stack.len()
	if exp != res {
		t.Errorf("expected:%d but got instead:%d", exp, res)
	}

	stack.push("1").push("23")
	exp = 2
	res = stack.len()
	if exp != res {
		t.Errorf("expected:%d but got instead:%d", exp, res)
	}
}
