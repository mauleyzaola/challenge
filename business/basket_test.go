package business

import (
	"strings"
	"testing"

	"github.com/mauleyzaola/challenge/domain"
)

// TODO the product price should come from another domain entity, for now we store it along with the product itself
func TestBasketAmount(t *testing.T) {
	voucher := &domain.Product{
		Code:  "VOUCHER",
		Price: 5,
	}
	tShirt := &domain.Product{
		Code:  "TSHIRT",
		Price: 20,
	}
	mug := &domain.Product{
		Code:  "MUG",
		Price: 7.5,
	}

	products := []domain.Product{*voucher, *tShirt, *mug}

	cases := []struct {
		rules    []domain.Rule
		products []domain.Product
		codes    string
		expected float64
		error    bool
	}{
		{
			rules:    domain.MockedRules,
			products: products,
			codes:    "VOUCHER,TSHIRT,MUG",
			expected: 32.5,
			error:    false,
		},
		{
			rules:    domain.MockedRules,
			products: products,
			codes:    "VOUCHER,TSHIRT,VOUCHER",
			expected: 25,
			error:    false,
		},
		{
			rules:    domain.MockedRules,
			products: products,
			codes:    "TSHIRT,TSHIRT,TSHIRT,VOUCHER,TSHIRT",
			expected: 81,
			error:    false,
		},
		{
			rules:    domain.MockedRules,
			products: products,
			codes:    "VOUCHER,TSHIRT,VOUCHER,VOUCHER,MUG,TSHIRT,TSHIRT",
			expected: 74.5,
			error:    false,
		},
		{
			rules:    domain.MockedRules,
			products: products,
			codes:    "SHA-LALALA",
			expected: 0,
			error:    true,
		},
	}

	for i, tc := range cases {
		items, err := BasketAddCode(strings.Split(tc.codes, ","), tc.products)
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
		result, err := BasketAmount(items, tc.rules)
		if result != tc.expected {
			t.Errorf("expected:%v but got instead:%v", tc.expected, result)
		}
	}
}
