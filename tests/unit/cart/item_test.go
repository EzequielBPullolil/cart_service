package cart_test

import (
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	"github.com/stretchr/testify/assert"
)

func TestItemEntity(t *testing.T) {

	t.Run("Should be invalid item update instance", func(t *testing.T) {
		var cases = []struct {
			Title string
			Item  cart.Item
		}{
			{"Because the name is empty", cart.Item{
				Name:     "",
				Price:    10,
				Quantity: 1,
				Currency: "ARS",
			}},
			{"Because the currency is empty", cart.Item{
				Name:     "A name",
				Price:    10,
				Quantity: 1,
				Currency: "",
			}},
			{"Because the cuantity is lower than 0", cart.Item{
				Name:     "A name",
				Price:    10,
				Quantity: -1,
				Currency: "ARS",
			}},
			{"Because the cuantity is equal 0", cart.Item{
				Name:     "A name",
				Price:    10,
				Quantity: 0,
				Currency: "ARS",
			}},
			{"Because the price is lower than 0", cart.Item{
				Name:     "A name",
				Price:    -10,
				Quantity: 10,
				Currency: "ARS",
			}},
			{"Because the price is equal 0", cart.Item{
				Name:     "A name",
				Price:    0,
				Quantity: 10,
				Currency: "ARS",
			}},
		}

		for _, c := range cases {
			t.Run(c.Title, func(t *testing.T) {
				assert.False(t, c.Item.ValidateUpdateFields())
			})
		}
	})

	t.Run("Should be a valid update_item", func(t *testing.T) {
		assert.True(t, cart.Item{
			Name:     "Abddd",
			Price:    float64(10),
			Quantity: 2,
			Currency: "ARS",
		}.ValidateUpdateFields())
	})
}
