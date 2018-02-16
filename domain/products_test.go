package domain

import (
	"reflect"
	"testing"
)

func TestFilterProducts(t *testing.T) {
	type tcase struct {
		input, expected Products
		cb              func(Product) bool
	}

	priceGt15 := func(p Product) bool {
		return p.Price > 15
	}

	testCases := []tcase{
		{
			input:    []Product{},
			expected: []Product{},
		},
		{
			input: []Product{
				{Code: "1", Name: "One", Price: 10},
				{Code: "2", Name: "Two", Price: 20},
			},
			expected: []Product{
				{Code: "2", Name: "Two", Price: 20},
			},
			cb: priceGt15,
		},
	}

	for _, tc := range testCases {
		results := tc.input.Filter(tc.cb)
		if !reflect.DeepEqual(tc.expected, results) {
			t.Errorf("expected:\n%#v\nbut got instead:\n%#v", tc.expected, results)
		}
	}
}
