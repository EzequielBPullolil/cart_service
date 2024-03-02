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

// modifies the fields and only the fields passed by parameter
// If any of the fields are not valid, do not modify any of them
func (item *Item) ModifieFields(new_fields Item) {
	if !new_fields.ValidateUpdateFields() {
		return
	}
	if new_fields.Currency != item.Currency {
		item.Currency = new_fields.Currency
	}

	if new_fields.Name != item.Name {
		item.Name = new_fields.Name
	}

	if new_fields.Quantity != item.Quantity {
		item.Quantity = new_fields.Quantity
	}

	if new_fields.Price != item.Price {
		item.Price = new_fields.Price
	}
}

func (item Item) ValidateUpdateFields() bool {

	return item.ValidateCurrency() && item.ValidateName() && item.ValidateQuantity() && item.ValidatePrice()
}

func (i Item) ValidateCurrency() bool {
	return i.Currency != ""
}
func (i Item) ValidateName() bool {
	return i.Name != ""
}

func (i Item) ValidateQuantity() bool {
	return i.Quantity > 0
}

func (i Item) ValidatePrice() bool {
	return i.Price > 0
}
