package main

import (
	"github.com/mauleyzaola/challenge/business"
	"github.com/mauleyzaola/challenge/domain"
	"github.com/mauleyzaola/challenge/operations"
	"github.com/mauleyzaola/challenge/storage"
)

type context struct {
	storage  operations.Storage
	products []domain.Product
	rules    []domain.Rule
}

func newContext() *context {
	ctx := &context{}

	ctx.storage = &storage.Memory{}
	ctx.storage.Init()

	// TODO load these rules from somewhere else
	ctx.rules = domain.MockedRules

	// TODO load these products from somewhere else
	ctx.products = []domain.Product{
		{Code: "VOUCHER", Price: 5, Name: "Cabify Voucher"},
		{Code: "TSHIRT", Price: 20, Name: "Cabify T-Shirt"},
		{Code: "MUG", Price: 7.5, Name: "Cafify Coffee Mug"},
	}

	return ctx
}

func (this *context) CreateBasket() (string, error) {
	return this.storage.Create()
}

func (this *context) ScanProduct(id string, codes []string) (domain.Products, error) {
	basket, err := this.storage.Load(id)
	if err != nil {
		return nil, err
	}
	items, err := business.BasketAddCode(codes, this.products)
	if err != nil {
		return nil, err
	}
	basket.Items = append(basket.Items, items...)
	if err = this.storage.Save(basket); err != nil {
		return nil, err
	}
	return items.DistinctProducts(), nil
}
