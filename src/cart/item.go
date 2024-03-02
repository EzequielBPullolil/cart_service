package cart

import (
	"github.com/google/uuid"
)

type Item struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

func CreateItem(name, currency string, price float64) *Item {
	return &Item{
		Id:       uuid.New().String(),
		Name:     name,
		Price:    price,
		Currency: currency,
	}
}
