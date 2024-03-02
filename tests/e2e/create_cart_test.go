package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	dbmanager "github.com/EzequielBPullolil/cart_service/src/db_manager"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateCart(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/carts", bytes.NewBufferString(`{
		"user": {
			"id": "fake user id"
		},
		"currency": "ARS"
	}`))
	app.ServeHTTP(w, req)

	var Response struct {
		Status string    `json:"status"`
		Data   cart.Cart `json:"data"`
	}
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&Response))
	t.Run("Response should", func(t *testing.T) {
		log.Println(w.Body)
		t.Run("Have cart data and status 201", func(t *testing.T) {
			assert.Equal(t, 201, w.Result().StatusCode)
			assert.NotNil(t, Response.Data)
		})

		t.Run("Have a cart with 0 items and 0 amount", func(t *testing.T) {
			cart := Response.Data

			assert.Empty(t, cart.Items)
			assert.Equal(t, float64(0), cart.Amount)
			assert.NoError(t, uuid.Validate(cart.Id))
		})
	})

	t.Run("Should persist cart", func(t *testing.T) {
		collection := dbmanager.ConnectDB(os.Getenv("DB_URI"), os.Getenv("DB_NAME")).CartCollection

		result := collection.FindOne(context.Background(), bson.M{"id": Response.Data.Id})

		assert.NotNil(t, result)
	})
}
