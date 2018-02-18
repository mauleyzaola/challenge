package domain

import "fmt"

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

func (this Products) MatchIndexes(indexes map[int]bool) Products {
	if indexes == nil {
		return this
	}
	var result []Product
	for i, v := range this {
		if _, ok := indexes[i]; ok {
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
