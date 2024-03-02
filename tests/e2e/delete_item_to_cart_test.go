package e2e

import (
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

func TestDeleteItemFromCart(t *testing.T) {
	cart_suject := cart.CreateCart("ARS")
	assert.NoError(t, cart_suject.Persist())

	item_suject := cart.CreateItem("an item", "ARS", 0.0, 1)
	assert.NoError(t, cart_suject.AddItemAndSave(*item_suject))
	t.Run("Should response with status 400 if The cart don't exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/carts/fake_cart/items/"+item_suject.Id, nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("Should response with status code 200 if the item does not exist in the cart", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/carts/"+cart_suject.Id+"/items/nonExistItem", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("If the item exist in the cart", func(t *testing.T) {

		items_before_change := cart_suject.Items
		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/carts/"+cart_suject.Id+"/items/"+item_suject.Id, nil)
		app.ServeHTTP(w, req)
		t.Run("Should remove the item from the cart in DB", func(t *testing.T) {
			var cart_from_db cart.Cart
			c := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection
			assert.NoError(t, c.FindOne(context.Background(), bson.M{"id": cart_suject.Id}).Decode(&cart_from_db))
			assert.NotEqual(t, cart_from_db.Items, items_before_change)
			assert.NotContains(t, cart_from_db.Items, item_suject)

		})

		t.Run("Should response with the updated cart", func(t *testing.T) {
			var response struct {
				Status string    `json:"status"`
				Cart   cart.Cart `json:"cart"`
			}
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
			assert.Equal(t, "cart updated", response.Status)
		})
	})
}
