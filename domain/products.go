package domain

import (
	"fmt"
)

type Products []Product

func (this Products) ToMap() (map[string]*Product, error) {
	result := make(map[string]*Product)
	for i := range this {
		v := this[i]
		if len(v.Code) == 0 {
			return nil, fmt.Errorf("missing code for product, cannot convert to map")
		}
		result[v.Code] = &v
	}
	return result, nil
}
