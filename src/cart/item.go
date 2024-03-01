package cart

import (
	"github.com/google/uuid"
)

type Item struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func CreateItem(name, price string) *Item {
	return &Item{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
