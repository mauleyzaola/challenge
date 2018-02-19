package business

import (
	"fmt"

	"github.com/mauleyzaola/challenge/domain"
)

func BasketAmount(codes []string, products domain.Products, rules []domain.Rule) (float64, error) {
	var result, amount float64

	// make sure all the provided codes are found in the products
	products, err := products.Distinct(codes)
	if err != nil {
		return 0, err
	}

	// iterate over each product and the first rule that matches will be applied for the discount
	// TODO check if more than one rule applies for each product
	for _, product := range products {
		var matchedRule *domain.Rule
		code := product.Code
		for i, rule := range rules {
			if rule.Code == code {
				matchedRule = &rules[i]
			}
		}
		tmp := products.Filter(func(p domain.Product) bool {
			return p.Code == code
		})
		if matchedRule != nil {
			callback, err := WhenParser(matchedRule.When, matchedRule.PriceExpr)
			if err != nil {
				return 0, err
			}
			amount, err = callback(tmp)
			if err != nil {
				return 0, err
			}
		} else {
			amount = float64(len(tmp)) * product.Price
		}
		fmt.Println("xxx:", amount, result, product.Code)
		result += amount

	}

	return result, nil
}
