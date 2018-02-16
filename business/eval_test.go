package business

import "testing"

func TestCompute(t *testing.T) {
	t.Skip()
	//type tcase struct {
	//	inputPrice, expectedPrice float64
	//	expr                      string
	//	error                     bool
	//}
	//cases := []tcase{
	//	{
	//		inputPrice:    5,
	//		expectedPrice: 0,
	//		expr:          "0",
	//	},
	//}
	//for _, tc := range cases {
	//	fn, err := evaluate(tc.expr)
	//	if tc.error {
	//		if err == nil {
	//			t.Errorf("expected error but got instead nil with expr:%s", tc.expr)
	//		}
	//	} else {
	//		if err != nil {
	//			t.Error(err)
	//			continue
	//		}
	//		result := fn(tc.inputPrice)
	//		if result != tc.expectedPrice {
	//			t.Errorf("expected price is:%v but got instead:%v", tc.expectedPrice, result)
	//		}
	//	}
	//}
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
	//{
	//	input:    "0",
	//	expected: 0,
	//},
	//{
	//	input:    "2+3",
	//	expected: 5,
	//},
	//{
	//	input:    "2-3",
	//	expected: -1,
	//},
	//{
	//	input:    "2*-3",
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
		result, err := evaluate(tc.input, tc.constants)
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
