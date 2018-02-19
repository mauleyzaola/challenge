package domain

import (
	"fmt"
)

type Products []Product

func (this Products) Filter(cb func(Product) bool) Products {
	if cb == nil {
		return this
	}
	var result []Product
	for _, v := range this {
		if cb(v) {
			result = append(result, v)
		}
	}
	return result
}

func (this Products) ToMap() (map[string]*Product, error) {
	result := make(map[string]*Product)
	for i := range this {
		v := this[i]
		if len(v.Code) == 0 {
			return nil, fmt.Errorf("missing code for product, cannot convert to map")
		}
		if _, ok := result[v.Code]; ok {
			return nil, fmt.Errorf("found duplicated code:%s", v.Code)
		}
		result[v.Code] = &v
	}
	return result, nil
}

func (this Products) ToProducts(codes []string) (Products, error) {
	var result Products
	keys := make(map[string]*Product)
	for i := range this {
		v := &this[i]
		keys[v.Code] = v
	}
	for _, code := range codes {
		product, ok := keys[code]
		if !ok {
			return nil, fmt.Errorf("code not found:%s", code)
		}
		result = append(result, *product)
	}
	return result, nil
}

func (this Products) Group() map[string]int {
	result := make(map[string]int)
	for _, v := range this {
		val, ok := result[v.Code]
		if !ok {
			val = 0
		}
		val++
		result[v.Code] = val
	}
	return result
}
