package domain

// Based on the content of a given slice of products, returns those positions of the slice which match some rule
type WhenCallback func(products []Product) (matches map[int]bool)
