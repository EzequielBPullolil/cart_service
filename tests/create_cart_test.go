package tests

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/carts", nil)
	app.ServeHTTP(w, req)

	t.Run("Response should", func(t *testing.T) {
		var Response struct {
			Status string    `json:"status"`
			Data   cart.Cart `json:"data"`
		}

		assert.NoError(t, json.NewDecoder(w.Body).Decode(&Response))
		log.Println(Response)
		t.Run("Have cart data and status 201", func(t *testing.T) {
			assert.Equal(t, 201, w.Result().StatusCode)
			assert.NotNil(t, Response.Data)
		})

		t.Run("Have a cart with 0 items and 0 amount", func(t *testing.T) {
			cart := Response.Data

			assert.Empty(t, cart.Items)
			assert.Equal(t, "0ARS", cart.Amount)
			assert.NoError(t, uuid.Validate(cart.Id))
		})
	})
}
