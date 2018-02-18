package domain

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
