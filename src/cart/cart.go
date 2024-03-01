package cart

import (
	"context"
	"os"

	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"github.com/google/uuid"
)

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
	cart_collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection

	_, err := cart_collection.InsertOne(context.Background(), c)

	return err
}

func FindCartById(id string) *Cart {
	var cart Cart
	cart_collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection

	result := cart_collection.FindOne(context.Background(), Cart{Id: id})

	result.Decode(&cart)

	return &cart
}
