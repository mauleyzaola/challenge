package domain

import "fmt"

type BasketItems []BasketItem

func (this BasketItems) Len() int {
	return len(this)
}

func (this BasketItems) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this BasketItems) Less(i, j int) bool {
	return this[i].Product.Code < this[j].Product.Code
}

func (this BasketItems) Filter(fn func(item BasketItem) bool) BasketItems {
	if fn == nil {
		return this
	}
	var result []BasketItem
	for _, v := range this {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (this BasketItems) CountProducts() map[string]int {
	result := make(map[string]int)
	for _, v := range this {
		val, ok := result[v.Product.Code]
		if !ok {
			val = 0
		}
		val += v.Quantity
		result[v.Product.Code] = val
	}
	return result
}

func (this BasketItems) ToItems(codes []string, products Products) (BasketItems, error) {
	prCodes, err := products.ToMap()
	if err != nil {
		return nil, err
	}
	var result BasketItems

	for _, code := range codes {
		product, ok := prCodes[code]
		if !ok {
			return nil, fmt.Errorf("code:%s was not found in provided products", code)
		}
		item := BasketItem{Product: product, Quantity: 1}
		result = append(result, item)
	}

	return result, nil
}
