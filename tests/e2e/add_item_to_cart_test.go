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
	cart_suject.Persist()
	assert.NoError(t, err)
	t.Run("Should response with status code 400", func(t *testing.T) {
		t.Run("if the cart dont exist", func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/carts/fake_id/items", bytes.NewBuffer(itemJSON))
			app.ServeHTTP(w, req)
			assert.Equal(t, 400, w.Result().StatusCode)
		})
		t.Run("if the item is invalid", func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/carts/"+cart_suject.Id+"/items", nil)
			app.ServeHTTP(w, req)
			assert.Equal(t, 400, w.Result().StatusCode)
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
