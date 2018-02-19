package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestProducts_Filter(t *testing.T) {
	priceGt15 := func(p Product) bool {
		return p.Price > 15
	}

	cases := []struct {
		input, expected Products
		callBack        func(Product) bool
	}{
		{
			input:    Products{},
			expected: Products{},
		},
		{
			input: Products{
				{Code: "1", Name: "One", Price: 10},
				{Code: "2", Name: "Two", Price: 20},
			},
			expected: Products{
				{Code: "2", Name: "Two", Price: 20},
			},
			callBack: priceGt15,
		},
	}

	for _, tc := range cases {
		results := tc.input.Filter(tc.callBack)
		if !reflect.DeepEqual(tc.expected, results) {
			t.Errorf("expected:\n%#v\nbut got instead:\n%#v", tc.expected, results)
		}
	}
}

func TestProducts_ToMap(t *testing.T) {
	cases := []struct {
		products Products
		expected map[string]*Product
		error    bool
	}{
		{
			products: Products{
				{Code: "1", Price: 10},
				{Code: "2", Price: 20},
				{Code: "3", Price: 30},
			},
			expected: map[string]*Product{
				"1": &Product{Code: "1", Price: 10},
				"2": &Product{Code: "2", Price: 20},
				"3": &Product{Code: "3", Price: 30},
			},
		},
	}

	marshal := func(v map[string]*Product) string {
		data, _ := json.Marshal(&v)
		return string(data)
	}

	for _, tc := range cases {
		result, err := tc.products.ToMap()
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil instead with products:%#v", tc.products)
			}
			continue
		}
		if err != nil {
			t.Error(err)
			continue
		}
		m1 := marshal(tc.expected)
		m2 := marshal(result)
		if m1 != m2 {
			t.Errorf("expected:\n%s\n but got instead:\n%s\n", m1, m2)
		}
	}
}
