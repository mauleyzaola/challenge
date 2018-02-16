package business

import "testing"

func TestStackCast(t *testing.T) {
	stack := &numberStack{}
	stack.Push("1").Push("2")
	expected := "12"
	result := stack.Value()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
	var fexpected float64 = 12
	fresult := stack.Float()
	if fexpected != fresult {
		t.Errorf("expected:%v but got instead:%v", fexpected, fresult)
	}
}

func TestStackReverse(t *testing.T) {
	stack := &numberStack{}
	stack.Push("123").Push("456")
	expected := "456123"
	result := stack.Reverse()
	if expected != result {
		t.Errorf("expected:%s but got instead:%s", expected, result)
	}
}

func TestStackSlice(t *testing.T) {
	t.Skip()
	stack := &numberStack{}
	stack.Push("1").Push("2")
	expected := "2"
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
}
