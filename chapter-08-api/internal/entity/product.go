package entity

import (
	"errors"
	"time"

	"github.com/lucassodresa/golang-post-graduation/chapter-08-api/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("iD is required")
	ErrInvalidID       = errors.New("invalid ID")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (product *Product) Validate() error {
	if product.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entity.ParseID(product.ID.String()); err != nil {
		return ErrInvalidID
	}

	if product.Name == "" {
		return ErrNameIsRequired
	}

	if product.Price == 0 {
		return ErrPriceIsRequired
	}

	if product.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
