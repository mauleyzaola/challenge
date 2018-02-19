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
