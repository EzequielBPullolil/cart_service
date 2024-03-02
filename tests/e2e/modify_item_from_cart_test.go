package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestModifyItemFromCart(t *testing.T) {
	cart_suject := cart.CreateCart("ARS")
	assert.NoError(t, cart_suject.Persist())
	item_suject := cart.CreateItem("an item", "ARS", 0.0, 1)
	itemJson, err := json.Marshal(item_suject)
	assert.NoError(t, err)
	assert.NoError(t, cart_suject.AddItemAndSave(*item_suject))

	t.Run("Should response status code 400 if", func(t *testing.T) {
		t.Run(" if The cart don't exist", func(t *testing.T) {
			var errorResponse ErrorResponse
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "/carts/fake_cart/items/"+item_suject.Id, bytes.NewBuffer(itemJson))

			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&errorResponse))
			assert.Equal(t, 400, w.Result().StatusCode)
			assert.Equal(t, "error modify item from cart", errorResponse.Status)
			assert.Equal(t, "Cart not found", errorResponse.Error)
		})
		t.Run(" if The item don't exist in the cart", func(t *testing.T) {
			var errorResponse ErrorResponse
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "/carts/"+cart_suject.Id+"/items/fakeUpdateItem", bytes.NewBuffer(itemJson))

			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&errorResponse))
			assert.Equal(t, 400, w.Result().StatusCode)
			assert.Equal(t, "error modify item from cart", errorResponse.Status)
			assert.Equal(t, "item dont exist in the cart", errorResponse.Error)
		})
		t.Run(" if The item is invalid", func(t *testing.T) {
			var errorResponse ErrorResponse
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "/carts/"+cart_suject.Id+"/items/"+item_suject.Id, nil)

			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&errorResponse))
			assert.Equal(t, 400, w.Result().StatusCode)
			assert.Equal(t, "error modify item from cart", errorResponse.Status)
			assert.Equal(t, "invalid item", errorResponse.Error)
		})
	})

	t.Run("Should response with status code 200", func(t *testing.T) {
		w := httptest.NewRecorder()
		updateJSON, err := json.Marshal(cart.Item{
			Name:     "new name",
			Price:    float64(9),
			Quantity: 10,
			Currency: "COL",
		})
		assert.NoError(t, err)
		req := httptest.NewRequest("PATCH", "/carts/"+cart_suject.Id+"/items/"+item_suject.Id, bytes.NewBuffer(updateJSON))
		app.ServeHTTP(w, req)
		t.Run("Should be modified in the db", func(t *testing.T) {
			var finded cart.Cart
			collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
			assert.NoError(t, collection.FindOne(context.Background(), bson.M{"id": cart_suject.Id}).Decode(&finded))
			assert.Contains(t, finded.Items, cart.Item{
				Id:       item_suject.Id,
				Name:     "new name",
				Price:    float64(9),
				Quantity: 10,
				Currency: "COL",
			})

		})
		t.Run("Should response with the expected data", func(t *testing.T) {
			var Respose struct {
				Status string    `json:"status"`
				Cart   cart.Cart `json:"cart"`
			}
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&Respose))
			assert.Equal(t, 200, w.Result().StatusCode)
			assert.Equal(t, "cart item updated", Respose.Status)
			assert.NotEqual(t, Respose.Cart.Items, cart_suject.Items)
		})
	})
}
