package domain

type DiscountRule struct {
	Code, Rule, Apply string
}

// Defines which discount rules would apply to which products, currently mocked in code, but could be stored anywhere else
//
// The idea is that logic is driven by data
var productDiscMap []DiscountRule = []DiscountRule{
	{
		Code:  "VOUCHER",
		Rule:  "each:2",
		Apply: "price=0",
	},
	{
		Code:  "TSHIRT",
		Rule:  "greater:3",
		Apply: "price*.95",
	},
}
