package cart

import (
	"github.com/google/uuid"
)

type Item struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Quantity int     `json:"quantity"`
}

// quantity must be greater than zero
func CreateItem(name, currency string, price float64, quantity int) *Item {
	if quantity <= 0 {
		return nil
	}
	return &Item{
		Id:       uuid.New().String(),
		Name:     name,
		Price:    price,
		Currency: currency,
		Quantity: quantity,
	}
}

func (item Item) CalculateAmount() float64 {
	return item.Price * float64(item.Quantity)
}
