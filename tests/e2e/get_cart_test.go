package e2e

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	"github.com/stretchr/testify/assert"
)

func TestGetCartFromDb(t *testing.T) {

	type Response struct {
		Cart cart.Cart `json:"cart"`
	}

	t.Run("Should return a cart if the id is registered", func(t *testing.T) {
		var response Response
		w := httptest.NewRecorder()
		cart := cart.CreateCart("ARS")
		assert.NoError(t, cart.Persist())
		req := httptest.NewRequest("GET", "/carts/"+cart.Id, nil)
		app.ServeHTTP(w, req)
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))

	})

	t.Run("Should response with empty json if the cart dont exist", func(t *testing.T) {
		var response Response
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/carts/fake_id", nil)
		app.ServeHTTP(w, req)
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
		assert.Equal(t, 200, w.Result().StatusCode)

		log.Println(response)
		assert.Empty(t, response.Cart)
	})

	t.Run("Should calculate the amount of cart", func(t *testing.T) {
		var cart_suject = cart.CreateCart("ARS")
		w := httptest.NewRecorder()
		assert.NoError(t, cart_suject.Persist())
		req := httptest.NewRequest("GET", "/carts/"+cart_suject.Id, nil)
		app.ServeHTTP(w, req)
		t.Run("Should be 0 if the cart dont have items", func(t *testing.T) {
			var response Response
			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
			assert.Equal(t, float64(0), response.Cart.Amount)
		})
		t.Run("Should not be 0 if the cart have items", func(t *testing.T) {
			var response Response
			expected_price := float64(20)
			item := cart.CreateItem("fake_item", "ARS", expected_price, 1)
			cart_suject.AddItemAndSave(*item)
			app.ServeHTTP(w, req)
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
			assert.NotEqual(t, 0, response.Cart.Amount)
			assert.Equal(t, expected_price, response.Cart.Amount)
		})
	})
}
