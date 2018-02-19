package operations

import (
	"github.com/mauleyzaola/challenge/domain"
)

type Storage interface {
	Init()

	List() []string
	Create() (*domain.Basket, error)
	Load(id string) (*domain.Basket, error)
	Save(basket *domain.Basket) error
	Remove(id string) error
}
