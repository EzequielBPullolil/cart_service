package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	"github.com/stretchr/testify/assert"
)

func TestGetCartFromDb(t *testing.T) {

	var Response struct {
		Cart cart.Cart `json:"cart"`
	}

	t.Run("Should return a cart if the id is registered", func(t *testing.T) {
		w := httptest.NewRecorder()
		cart := cart.CreateCart()
		assert.NoError(t, cart.Persist())
		req := httptest.NewRequest("GET", "/carts/"+cart.Id, nil)
		app.ServeHTTP(w, req)
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&Response))

	})

	t.Run("Should response with empty json if the cart dont exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/carts/fake_id", nil)
		app.ServeHTTP(w, req)
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&Response))
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.Empty(t, Response)
	})
}
