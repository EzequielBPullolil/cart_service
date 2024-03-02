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

func TestAddItemToCart(t *testing.T) {
	var item = cart.CreateItem("fake name", "ARS", float64(20), 1)
	itemJSON, err := json.Marshal(item)
	cart_suject := cart.CreateCart("ARS")
	assert.NoError(t, cart_suject.Persist())
	assert.NoError(t, err)
	t.Run("Should response with status code 400", func(t *testing.T) {
		t.Run("if the cart dont exist", func(t *testing.T) {
			var errorResponse ErrorResponse
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/carts/fake_id/items", bytes.NewBuffer(itemJSON))
			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&errorResponse))
			assert.Equal(t, 400, w.Result().StatusCode)
			assert.Equal(t, errorResponse.Status, "error adding item to cart")
			assert.Equal(t, errorResponse.Error, "Cart not found")
		})
		t.Run("if the item is invalid", func(t *testing.T) {
			var errorResponse ErrorResponse
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/carts/"+cart_suject.Id+"/items", nil)
			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&errorResponse))
			assert.Equal(t, 400, w.Result().StatusCode)

			assert.Equal(t, errorResponse.Status, "error adding item to cart")
			assert.Equal(t, errorResponse.Error, "invalid item")
		})

		t.Run("If the item already exist in the cart", func(t *testing.T) {
			var errorResponse ErrorResponse
			fail_cart := cart.CreateCart("ARS")
			assert.NoError(t, fail_cart.Persist())
			item_ := cart.CreateItem("alreadyRegistered item", "ARS", 0, 1)
			json_item, err := json.Marshal(item_)
			assert.NoError(t, err)
			assert.NoError(t, fail_cart.AddItemAndSave(*item_))
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/carts/"+fail_cart.Id+"/items", bytes.NewBuffer(json_item))
			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&errorResponse))
			assert.Equal(t, 400, w.Result().StatusCode)
			assert.Equal(t, "error adding item to cart", errorResponse.Status)
			assert.Equal(t, "the item is already in the cart", errorResponse.Error)
		})
	})

	t.Run("Should response with status code 201 and the updated cart", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/carts/"+cart_suject.Id+"/items", bytes.NewBuffer(itemJSON))
		app.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Result().StatusCode)
		t.Run("Should be persisted in DB", func(t *testing.T) {
			var finded cart.Cart
			collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
			collection.FindOne(context.Background(), bson.M{"id": cart_suject.Id}).Decode(&finded)

			assert.Equal(t, finded.Items[0].Name, item.Name)
			assert.Equal(t, finded.Items[0].Id, item.Id)
		})
	})

}
