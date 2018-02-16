package domain

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
