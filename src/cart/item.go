package cart

import (
	"github.com/google/uuid"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func CreateItem(name string) *Item {
	return &Item{
		Id:   uuid.New().String(),
		Name: name,
	}
}
