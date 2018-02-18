package business

import (
	"reflect"
	"strings"
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

func TestOpOrder(t *testing.T) {
	cases := []struct {
		op       string
		expected int
	}{
		{op: "+", expected: 1},
		{op: "-", expected: 1},
		{op: "*", expected: 2},
		{op: "/", expected: 2},
		{op: "", expected: 0},
		{op: "x", expected: 0},
	}
	for _, tc := range cases {
		result := operatorOrder(tc.op)
		if result != tc.expected {
			t.Errorf("expected:%v but got instead:%v op:%s", tc.expected, result, tc.op)
		}
	}
}

func TestInfixToPostfix(t *testing.T) {
	cases := []struct {
		input, expected string
		error           bool
		constants       map[string]float64
	}{
		{
			input:    "3+4*2/(1-5)",
			expected: "3,4,2,*,1,5,-,/,+",
			error:    false,
		},
		{
			input:    "3+44*2/(1-5)",
			expected: "3,44,2,*,1,5,-,/,+",
			error:    false,
		},
		{
			input:    "3+fortyFour*2/(1-5)",
			expected: "3,44.000000,2,*,1,5,-,/,+",
			error:    false,
			constants: map[string]float64{
				"fortyFour": 44,
			},
		},
		{
			input:    "3+44.88*22.01/(1-0.005)",
			expected: "3,44.88,22.01,*,1,0.005,-,/,+",
			error:    false,
		},
		{
			input:    "3+44.88*22.01/(1-0..005)",
			expected: "",
			error:    true,
		},
		{
			input:    "33lalala",
			expected: "",
			error:    true,
		},
	}
	for _, tc := range cases {
		result, err := infixToPostfix(tc.input, tc.constants)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil with input:%s", tc.input)
			}
		} else {
			if err != nil {
				t.Error(err)
			} else if tc.expected != strings.Join(result, ",") {
				t.Errorf("expected:%s but got instead:%s", tc.expected, result)
			}
		}
	}
}

func TestPostfixCalculator(t *testing.T) {
	cases := []struct {
		input    string
		expected float64
		error    bool
	}{
		{
			input:    "3,4,2,*,1,5,-,/,+",
			expected: 1,
			error:    false,
		},
		{
			input:    "3,44,2,*,1,5,-,/,+",
			expected: -19,
			error:    false,
		},
	}

	for _, tc := range cases {
		result, err := postfixCalculator(strings.Split(tc.input, ","))
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil with input:%s", tc.input)
			}
		} else {
			if err != nil {
				t.Error(err)
				continue
			}
			if tc.expected != result {
				t.Errorf("expected:%v but got instead:%v", tc.expected, result)
			}
		}
	}
}

func TestCalc(t *testing.T) {
	cases := []struct {
		input     string
		expected  float64
		error     bool
		constants map[string]float64
	}{
		{
			input:    "8+2*5",
			expected: 18,
			error:    false,
		},
		{
			input:    "4/(ten/ten5)+hundred-22",
			expected: 80,
			constants: map[string]float64{
				"ten":     10,
				"hundred": 100,
				"ten5":    5,
			},
			error: false,
		},
		{
			input:    "20*.95",
			expected: 19,
			error:    false,
		},
		{
			input:    "5-5",
			expected: 0,
			error:    false,
		},
		{
			input: "4/(10/5)-+100-22",
			error: true,
		},
	}

	for _, tc := range cases {
		values, err := infixToPostfix(tc.input, tc.constants)
		if !tc.error && err != nil {
			t.Errorf("expected error to be nil but got instead:%s with input:%s", err, tc.input)
			continue
		}
		result, err := postfixCalculator(values)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil instead with input:%s and postfix values:%s", tc.input, strings.Join(values, ","))
			}
			continue
		}
		if err != nil {
			t.Errorf("expected error to be nil but got instead:%s with input:%s", err, tc.input)
			continue
		}
		if result != tc.expected {
			t.Errorf("expected:%v but got instead:%v with input:%s", tc.expected, result, tc.input)
		}
	}
}

func TestSortBySize(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{
			input:    "price,prices,longlongname",
			expected: "longlongname,prices,price",
		},
	}
	for _, tc := range cases {
		slice := strings.Split(tc.input, ",")
		result := sortBySize(slice)
		if tc.expected != strings.Join(result, ",") {
			t.Errorf("expected:%s but got instead:%s", tc.expected, strings.Join(result, ","))
		}
	}
}

func TestSortConstants(t *testing.T) {
	cases := []struct {
		input, expected map[string]float64
	}{
		{
			input: map[string]float64{
				"price":     5,
				"lastPrice": 18.55,
				"x":         22,
			},
			expected: map[string]float64{
				"lastPrice": 18.55,
				"price":     5,
				"x":         22,
			},
		},
	}
	for _, tc := range cases {
		result := sortConstants(tc.input)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("expected:%#v but got instead:%#v", tc.expected, result)
		}
	}
}

func TestEval(t *testing.T) {
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
