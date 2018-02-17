package business

import "testing"

func TestToPostfix(t *testing.T) {
	input := "2+5*(3+8)"
	result := toPostfix(input)
	t.Log("xxx:", result)
}

func TestEvaluate(t *testing.T) {
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
