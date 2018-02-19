package business

import (
	"fmt"

	"github.com/mauleyzaola/challenge/domain"
)

func BasketAddCode(codes []string, products domain.Products) (domain.BasketItems, error) {
	var result domain.BasketItems
	prCodes, err := products.ToMap()
	if err != nil {
		return nil, err
	}

	counters := result.CountCodes(codes)
	for code, count := range counters {
		product, ok := prCodes[code]
		if !ok {
			return nil, fmt.Errorf("cannot find any code:%s", code)
		}
		item := domain.BasketItem{Product: product, Quantity: count}
		result = append(result, item)
	}
	return result, nil
}

func BasketAmount(items domain.BasketItems, rules []domain.Rule) (float64, error) {
	var result, amount float64

	// iterate over each product and the first rule that matches will be applied for the discount
	// TODO check if more than one rule applies for each product
	for _, item := range items {
		var matchedRule *domain.Rule
		for i, rule := range rules {
			if rule.Code == item.Product.Code {
				matchedRule = &rules[i]
			}
		}
		if matchedRule != nil {
			callback, err := WhenParser(matchedRule.When, matchedRule.PriceExpr)
			if err != nil {
				return 0, err
			}
			amount, err = callback(&item)
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
