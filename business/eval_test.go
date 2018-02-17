package business

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	cases := []struct {
		expected bool
		input    uint8
	}{
		{input: '1', expected: true},
		{input: 'x', expected: false},
	}
	for _, tc := range cases {
		result := isNumber(tc.input)
		if result != tc.expected {
			t.Errorf("expected:%v but got instead:%v input:%s", tc.expected, result, string(tc.input))
		}
	}
}

func TestIsOperator(t *testing.T) {
	cases := []struct {
		expected bool
		input    uint8
	}{
		{input: '+', expected: true},
		{input: '-', expected: true},
		{input: '*', expected: true},
		{input: '/', expected: true},
		{input: ' ', expected: false},
		{input: 'x', expected: false},
	}
	for _, tc := range cases {
		result := isOperator(tc.input)
		if result != tc.expected {
			t.Errorf("expected:%v but got instead:%v input:%s", tc.expected, result, string(tc.input))
		}
	}
}

func TestToPostfix(t *testing.T) {
	input := "3+4*2/(1-5)"
	result, err := infixToPostfix(input)
	if err != nil {
		t.Error(err)
	}
	t.Log("xxx:", result)
}

func TestEvaluate(t *testing.T) {
	t.Skip()
	type tcase struct {
		input     string
		constants map[string]float64
		expected  float64
		error     bool
	}
	cases := []tcase{
		//{
		//	input:    "price-price",
		//	expected: 0,
		//	constants: map[string]float64{
		//		"price": 5,
		//	},
		//},
		//{
		//	input:    "price*.95",
		//	expected: 19,
		//	constants: map[string]float64{
		//		"price": 20,
		//	},
		//},
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "2+3",
			expected: 5,
		},
		//{
		//	input:    "2-3",
		//	expected: -1,
		//},
		//{
		//	input:    "2*3",
		//	expected: -6,
		//},
		//{
		//	input:    "3/2",
		//	expected: 1.5,
		//},
		//{
		//	input:    "10+6-8*2",
		//	expected: 16,
		//},
		//{
		//	input: "",
		//	error: true,
		//},
		//{
		//	input: "--2",
		//	error: true,
		//},
	}
	for _, tc := range cases {
		result, err := eval(tc.input, tc.constants)
		if tc.error {

		} else {
			if err != nil {
				t.Error(err)
				continue
			}
			if result != tc.expected {
				t.Errorf("expected:%v but got instead:%v", tc.expected, result)
			}
		}
	}
}
