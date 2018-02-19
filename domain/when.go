package domain

// Callback that processes a basket item and returns the calculated price
type WhenCallback func(item *BasketItem) (float64, error)
