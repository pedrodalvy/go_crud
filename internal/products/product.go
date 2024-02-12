package products

import "github.com/google/uuid"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewProduct(name string) Product {
	return Product{
		ID:   uuid.New().String(),
		Name: name,
	}
}
