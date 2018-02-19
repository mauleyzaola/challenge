package domain

// Defines which discount rule to apply to certain code
type Rule struct {
	// The product code
	Code string

	// The matching rule to be checked against
	When string

	// The price transformation when the rule matches
	PriceExpr string
}
