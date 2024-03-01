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

func (c *Cart) AddItemAndSave(item_to_add Item) error {
	cart_collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
	if err := c.AddItem(item_to_add); err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"items": c.Items}}
	_, error := cart_collection.UpdateOne(context.Background(), bson.M{"id": c.Id}, update)
	return error
}

func (c *Cart) AddItem(item_to_add Item) error {
	for _, v := range c.Items {
		if v.Id == item_to_add.Id {
			return errors.New("the item is already in the cart")
		}
	}
	c.Items = append(c.Items, item_to_add)
	return nil
}
