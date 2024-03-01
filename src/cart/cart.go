package cart

import "github.com/google/uuid"

type Cart struct {
	Id     string     `json:"id"`
	Amount string     `json:"amount"`
	Items  []struct{} `json:"items"`
}

func CreateCart() *Cart {
	return &Cart{
		Id:     uuid.New().String(),
		Amount: "0ARS",
		Items:  make([]struct{}, 0),
	}
}

func (c *Cart) Persist() error {
	return nil
}
