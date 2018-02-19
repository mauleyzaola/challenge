package domain

type BasketItems []BasketItem

func (this BasketItems) CountCodes(codes []string) map[string]int {
	result := make(map[string]int)
	for _, code := range codes {
		val, ok := result[code]
		if !ok {
			val = 0
		}
		val++
		result[code] = val
	}
	return result
}

func (this BasketItems) DistinctProducts() Products {
	var result Products
	codes := make(map[string]struct{})
	for _, v := range this {
		_, ok := codes[v.Product.Code]
		if !ok {
			codes[v.Product.Code] = struct{}{}
			result = append(result, *v.Product)
		}
	}
	return result
}

func (this BasketItems) Group() BasketItems {
	var result BasketItems
	items := make(map[string]*BasketItem)
	for _, v := range this {
		val, ok := items[v.Product.Code]
		if !ok {
			val = &BasketItem{Product: v.Product}
			items[v.Product.Code] = val
		}
		val.Quantity += v.Quantity
	}
	for _, v := range items {
		result = append(result, *v)
	}
	return result
}
