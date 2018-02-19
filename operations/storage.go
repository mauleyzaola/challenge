package operations

import (
	"github.com/mauleyzaola/challenge/domain"
)

// TODO define lock and unlock methods
type Storage interface {
	List() []domain.Basket
	Load(id string) (*domain.Basket, error)
	Save(basket *domain.Basket) error
	Remove(id string) error
}
