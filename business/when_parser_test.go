package business

import (
	"strings"
	"testing"

	"github.com/mauleyzaola/challenge/domain"
)

func TestWhenEachParser(t *testing.T) {
	t.Skip()
	voucher := &domain.Product{Code: "VOUCHER", Name: "Cabify Voucher", Price: 5}
	tShirt := &domain.Product{Code: "TSHIRT", Name: "Cabify T-Shirt", Price: 20}
	mug := &domain.Product{Code: "MUG", Name: "Cafify Coffee Mug", Price: 7.5}

	products := []domain.Product{*voucher, *tShirt, *mug}
	cases := []struct {
		whenExpr, priceExpr string
		codes               []string
		products            domain.Products
		error               bool
		expected            float64
	}{
		{
			whenExpr: "",
			error:    true,
		},
		{
			whenExpr:  "each:2",
			products:  products,
			expected:  voucher.Price * 4 / 2,
			priceExpr: "price-price",
			codes:     strings.Split("VOUCHER,VOUCHER,VOUCHER,VOUCHER", ","),
		},
		{
			whenExpr:  "each:3",
			products:  products,
			expected:  voucher.Price * 3,
			priceExpr: "price-price",
			codes:     strings.Split("VOUCHER,VOUCHER,VOUCHER,VOUCHER", ","),
		},
	}

	for _, tc := range cases {
		callback, err := WhenParser(tc.whenExpr, tc.priceExpr)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil with whenExpr:%s", tc.whenExpr)
			}
			continue
		}
		if err != nil {
			t.Error(err)
			continue
		}
		if err != nil {
			t.Error("unexpected error:", err)
			continue
		}
		productCodes, err := tc.products.Distinct(tc.codes)
		if err != nil {
			t.Error("unexpected error:", err)
			continue
		}
		result, err := callback(productCodes)
		if err != nil {
			t.Error("unexpected error:", err)
			continue
		}
		if tc.expected != result {
			t.Errorf("expected:\n%#v\nbut got instead:\n%#v\n", tc.expected, result)
		}
	}
}

func TestWhenTotalCounter(t *testing.T) {
	t.Skip()
	voucher := &domain.Product{Code: "VOUCHER", Name: "Cabify Voucher", Price: 5}
	tShirt := &domain.Product{Code: "TSHIRT", Name: "Cabify T-Shirt", Price: 20}
	mug := &domain.Product{Code: "MUG", Name: "Cafify Coffee Mug", Price: 7.5}

	products := []domain.Product{*voucher, *tShirt, *mug}
	cases := []struct {
		whenExpr, priceExpr string
		codes               []string
		products            domain.Products
		error               bool
		expected            float64
	}{
		{
			whenExpr: "",
			error:    true,
		},
		{
			whenExpr:  "gt:3",
			products:  products,
			expected:  voucher.Price * 4 * .95,
			priceExpr: "price*.95",
			codes:     strings.Split("VOUCHER,VOUCHER,VOUCHER,VOUCHER", ","),
		},
		{
			whenExpr:  "gt:4",
			products:  products,
			expected:  (voucher.Price * 4) + (mug.Price * 1),
			priceExpr: "price*.95",
			codes:     strings.Split("VOUCHER,VOUCHER,MUG,VOUCHER,VOUCHER", ","),
		},
	}

	for _, tc := range cases {
		callback, err := WhenParser(tc.whenExpr, tc.priceExpr)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil with whenExpr:%s", tc.whenExpr)
			}
		} else {
			if err != nil {
				t.Error(err)
				continue
			}
			if err != nil {
				t.Error("unexpected error:", err)
				continue
			}
			productCodes, err := tc.products.Distinct(tc.codes)
			if err != nil {
				t.Error("unexpected error:", err)
				continue
			}
			result, err := callback(productCodes)

			if err != nil {
				t.Error("unexpected error:", err)
				continue
			}
			if tc.expected != result {
				t.Errorf("expected:\n%#v\nbut got instead:\n%#v\n", tc.expected, result)
			}
		}
	}
}
