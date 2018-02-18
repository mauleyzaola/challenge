package business

import (
	"testing"

	"github.com/mauleyzaola/challenge/domain"
)

// TODO the product price should come from another domain entity
func TestBasketAmount(t *testing.T) {
	t.Skip()
	voucher := &domain.Product{
		Code:  "VOUCHER",
		Price: 5,
	}
	tshirt := &domain.Product{
		Code:  "TSHIRT",
		Price: 20,
	}
	mug := &domain.Product{
		Code:  "MUG",
		Price: 7.5,
	}

	cases := []struct {
		rules    []domain.DiscountRule
		items    []domain.BasketItem
		expected float64
		error    bool
	}{
		{
			rules: domain.MockedDiscountRules,
			items: []domain.BasketItem{
				{Product: voucher, Quantity: 1},
				{Product: tshirt, Quantity: 1},
				{Product: mug, Quantity: 1},
			},
			expected: 0,
			error:    false,
		},
	}

	for i, tc := range cases {
		result, err := BasketAmount(tc.items, tc.rules)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got instead nil with test case:%d", i+1)
			}
			continue
		}
		if err != nil {
			t.Error("unexpected error:", err)
			continue
		}
		if result != tc.expected {
			t.Errorf("expected:%v but got instead:%v", tc.expected, result)
		}
	}
}
