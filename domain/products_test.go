package domain

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
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

func TestProducts_MatchIndexes(t *testing.T) {
	cases := []struct {
		products, expected Products
		indexes            map[int]bool
	}{
		{
			products: Products{},
			expected: Products{},
		},
		{
			products: Products{
				{Code: "1", Price: 1},
				{Code: "2", Price: 2},
				{Code: "3", Price: 3},
			},
			expected: Products{
				{Code: "1", Price: 1},
				{Code: "3", Price: 3},
			},
			indexes: map[int]bool{
				0: true,
				2: true,
			},
		},
		{
			products: Products{
				{Code: "1", Price: 1},
				{Code: "2", Price: 2},
				{Code: "3", Price: 3},
			},
			expected: Products{
				{Code: "3", Price: 3},
			},
			indexes: map[int]bool{
				2: true,
			},
		},
	}
	for _, tc := range cases {
		result := tc.products.MatchIndexes(tc.indexes)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("expected:%#v but got instead:%#v", tc.expected, result)
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

func TestProducts_ToItems(t *testing.T) {
	voucher := &Product{
		Code: "VOUCHER",
	}
	tshirt := &Product{
		Code: "TSHIRT",
	}
	mug := &Product{
		Code: "MUG",
	}
	products := Products{*voucher, *tshirt, *mug}

	cases := []struct {
		products Products
		codes    string
		expected BasketItems
		error    bool
	}{
		{
			products: products,
			codes:    "VOUCHER,TSHIRT,MUG",
			expected: BasketItems{
				{Product: mug, Quantity: 1},
				{Product: tshirt, Quantity: 1},
				{Product: voucher, Quantity: 1},
			},
			error: false,
		},
		{
			products: products,
			codes:    "VOUCHER,TSHIRT,VOUCHER",
			expected: BasketItems{
				{Product: voucher, Quantity: 2},
				{Product: tshirt, Quantity: 1},
			},
			error: false,
		},
		{
			products: products,
			codes:    "TSHIRT,TSHIRT,TSHIRT,VOUCHER,TSHIRT",
			expected: BasketItems{
				{Product: voucher, Quantity: 1},
				{Product: tshirt, Quantity: 4},
			},
			error: false,
		},
	}
	marshal := func(values []BasketItem) string {
		data, _ := json.Marshal(&values)
		return string(data)
	}
	for _, tc := range cases {
		result, err := tc.products.ToItems(strings.Split(tc.codes, ","))
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil instead with products:%#v", tc.products)
			}
			continue
		}
		if err != nil {
			t.Error("unexpected error:", err)
			continue
		}
		// avoid the slice to change its order when marshaling so we can compare correctly
		sort.Sort(tc.expected)
		sort.Sort(result)
		if marshal(tc.expected) != marshal(result) {
			t.Errorf("expected:\n%#v\nbut got instead:\n%#v\n", marshal(tc.expected), marshal(result))
		}
	}
}
