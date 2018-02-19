package domain

// Callback that processes a slice of products and returns the calculated price for each one
type WhenCallback func(codes []string, products Products) (float64, error)
