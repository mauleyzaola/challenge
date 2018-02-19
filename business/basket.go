package business

import (
	"fmt"

	"github.com/mauleyzaola/challenge/domain"
)

func BasketAmount(codes []string, products domain.Products, rules []domain.Rule) (float64, error) {
	var (
		result, amount float64
		items          domain.BasketItems
	)

	prCodes, err := products.ToMap()
	if err != nil {
		return 0, err
	}

	counters := items.CountCodes(codes)

	// iterate over each product and the first rule that matches will be applied for the discount
	// TODO check if more than one rule applies for each product
	for code, count := range counters {
		var matchedRule *domain.Rule
		for i, rule := range rules {
			if rule.Code == code {
				matchedRule = &rules[i]
			}
		}
		product, ok := prCodes[code]
		if !ok {
			return 0, fmt.Errorf("cannot find any code:%s", code)
		}
		item := &domain.BasketItem{Product: product, Quantity: count}
		if matchedRule != nil {
			callback, err := WhenParser(matchedRule.When, matchedRule.PriceExpr)
			if err != nil {
				return 0, err
			}
			amount, err = callback(item)
			if err != nil {
				return 0, err
			}
		} else {
			amount = float64(item.Quantity) * item.Product.Price
		}
		result += amount
	}

	return result, nil
}
