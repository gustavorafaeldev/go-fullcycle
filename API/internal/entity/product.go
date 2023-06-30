package entity

import (
	"errors"
	"go-curso/api/internal/entity"
	"go-curso/api/pkg/entity"
	"time"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrIdIsInvalid = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIsInvalid = errors.New("invalid price")
)

type Product struct {
	ID entity.ID `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIdIsInvalid
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrPriceIsInvalid
	}
	return nil
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}