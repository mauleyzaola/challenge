package domain

type DiscountRule struct {
	Code, When, Expr string
}

// Defines which discount rules would apply to which products, currently mocked in code, but could be stored anywhere else
//
// The idea is that logic is driven by data
var productDiscMap []DiscountRule = []DiscountRule{
	{
		Code: "VOUCHER",
		When: "each:2",
		Expr: "0",
	},
	{
		Code: "TSHIRT",
		When: "gte:3",
		Expr: "*.95",
	},
}
