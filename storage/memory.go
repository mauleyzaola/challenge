package storage

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/mauleyzaola/challenge/domain"
)

type Memory struct {
	sync.Mutex
	nextId int64
	items  map[string]interface{}
}

func (this *Memory) Init() {
	this.nextId = 0
	this.items = make(map[string]interface{})
}

func (this *Memory) List() []string {
	this.Lock()
	defer this.Unlock()
	var results []string
	for k := range this.items {
		results = append(results, k)
	}
	return results
}

func (this *Memory) Load(id string) (*domain.Basket, error) {
	this.Lock()
	defer this.Unlock()
	val, ok := this.items[id]
	if !ok {
		return nil, fmt.Errorf("no matching id found:%s", id)
	}
	basket, ok := val.(*domain.Basket)
	if !ok {
		return nil, fmt.Errorf("cannot cast to basket:%#v", basket)
	}
	return basket, nil
}

func (this *Memory) Create() (*domain.Basket, error) {
	this.Lock()
	defer this.Unlock()

	// mocked id just for tests
	this.nextId++
	id := strconv.FormatInt(this.nextId, 10)
	basket := &domain.Basket{
		Id:    id,
		Items: []domain.BasketItem{},
	}
	this.items[id] = basket
	return basket, nil
}

func (this *Memory) Save(basket *domain.Basket) error {
	this.Lock()
	defer this.Unlock()
	_, ok := this.items[basket.Id]
	if !ok {
		return fmt.Errorf("basket with id:%s does not exist", basket.Id)
	}
	this.items[basket.Id] = basket
	return nil
}

func (this *Memory) Remove(id string) error {
	this.Lock()
	defer this.Unlock()
	_, ok := this.items[id]
	if !ok {
		return fmt.Errorf("basket with id:%s does not exist", id)
	}
	delete(this.items, id)
	return nil
}
