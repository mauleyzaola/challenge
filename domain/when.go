package domain

// Callback that processes a slice of products and returns the calculated price for each one
type WhenCallback func(item *BasketItem) (float64, error)
