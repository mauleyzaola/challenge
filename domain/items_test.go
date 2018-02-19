package domain

import "testing"

func TestBasketItems_Group(t *testing.T) {
	products := []Product{
		{Code: "1", Name: "One", Price: 1},
		{Code: "2", Name: "Two", Price: 2},
		{Code: "3", Name: "Three", Price: 3},
	}
	var (
		items BasketItems
		total int
	)
	for i := 0; i < 100; i++ {
		p := &products[i%3]
		item := BasketItem{Product: p, Quantity: i}
		items = append(items, item)
		total += i
	}
	items = items.Group()
	expected := 3
	if len(items) != expected {
		t.Errorf("expected the items to be:%d but got instead:%d", expected, len(items))
	}

	expected = total
	total = 0
	for _, item := range items {
		total += item.Quantity
	}
	if expected != total {
		t.Errorf("expected the quantity to be:%d but got instead:%d", expected, total)
	}
}
