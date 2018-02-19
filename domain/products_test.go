package domain

import (
	"encoding/json"
	"testing"
)

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
