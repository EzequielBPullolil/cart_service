package cart

import (
	"context"
	"errors"
	"log"
	"os"

	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Cart struct {
	Id       string      `json:"id"`
	Amount   float64     `json:"amount"`
	Currency string      `json:"currency"`
	Items    []Item      `json:"items"`
	User     interface{} `json:"user"`
}

func CreateCart(currency string) *Cart {
	return &Cart{
		Id:       uuid.New().String(),
		Amount:   float64(0),
		Items:    make([]Item, 0),
		Currency: currency,
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

// Calculate cart amount
func (cart *Cart) CalculateAmount() {
	var amount float64
	for _, item := range cart.Items {
		amount += item.CalculateAmount()
	}
	cart.Amount = amount
}

// Add item to cart and persist in DB
func (c *Cart) AddItemAndSave(item_to_add Item) error {
	cart_collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
	if err := c.AddItem(item_to_add); err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"items": c.Items}}
	_, error := cart_collection.UpdateOne(context.Background(), bson.M{"id": c.Id}, update)
	return error
}

// Add item to items list if the item is not in the list
func (c *Cart) AddItem(item_to_add Item) error {
	for _, v := range c.Items {
		if v.Id == item_to_add.Id {
			return errors.New("the item is already in the cart")
		}
	}
	c.Items = append(c.Items, item_to_add)
	return nil
}

func (c *Cart) RemoveItemAndSave(item_id string) error {
	cart_collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
	c.RemoveItem(item_id)

	update := bson.M{"$set": bson.M{"items": c.Items}}
	_, error := cart_collection.UpdateOne(context.Background(), bson.M{"id": c.Id}, update)

	return error
}

// Change the cart list to one without the Item with the ID passed by parameter
// If there is no item with the passed Id then the list does not change
func (c *Cart) RemoveItem(item_id string) {
	new_items := make([]Item, 0)
	for _, v := range c.Items {
		if v.Id != item_id {
			new_items = append(new_items, v)
		}
	}

	c.Items = new_items
}

// modify the cart item with the passed id and update it in the database
func (c *Cart) ModifyItemAndSave(item_id string, new_fields Item) error {
	cart_collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
	if !c.ModifyItem(item_id, new_fields) {
		return errors.New("item dont exist in the cart")
	}

	update := bson.M{"$set": bson.M{"items": c.Items}}
	_, error := cart_collection.UpdateOne(context.Background(), bson.M{"id": c.Id}, update)

	return error
}

// modifies the cart item with the passed id and returns true if the item was found and modified
func (c *Cart) ModifyItem(item_id string, new_fields Item) bool {
	new_items := make([]Item, 0)
	found_and_modified := false
	for _, v := range c.Items {
		if v.Id == item_id {
			found_and_modified = true
			v.ModifieFields(new_fields)
		}
		new_items = append(new_items, v)
	}

	c.Items = new_items
	log.Println(c.Items)
	return found_and_modified
}
