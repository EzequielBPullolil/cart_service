package cart

import (
	"context"
	"errors"
	"os"

	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Cart struct {
	Id     string `json:"id"`
	Amount string `json:"amount"`
	Items  []Item `json:"items"`
}

func CreateCart() *Cart {
	return &Cart{
		Id:     uuid.New().String(),
		Amount: "0ARS",
		Items:  make([]Item, 0),
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

	err := cart_collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&cart)
	if err != nil {
		return nil
	}
	return &cart
}

	result.Decode(&cart)

	return &cart
}
