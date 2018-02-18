package domain

type DiscountRule struct {
	Code, When, PriceExpr string
}

// Defines which discount rules would apply to which products, currently mocked in code, but could be stored anywhere else
//
// The idea is that logic is driven by data and code each case separately once
var MockedDiscountRules []DiscountRule = []DiscountRule{
	{
		Code:      "VOUCHER",
		When:      "each:2",
		PriceExpr: "price-price",
	},
	{
		Code:      "TSHIRT",
		When:      "gte:3",
		PriceExpr: "price*.95",
	},
}
