package cart_test

import (
	"testing"

	"github.com/EzequielBPullolil/cart_service/src/cart"
	"github.com/stretchr/testify/assert"
)

func TestCart(t *testing.T) {

	t.Run("Should increase the amount as expected", func(t *testing.T) {
		cart_suject := cart.CreateCart("ARS")
		assert.Equal(t, 0.0, cart_suject.Amount)
		items := []cart.Item{
			*cart.CreateItem("cart1", "ARS", 20, 5),
			*cart.CreateItem("cart1", "ARS", 10, 1),
			*cart.CreateItem("cart1", "ARS", 9, 1),
			*cart.CreateItem("cart1", "ARS", 6, 1),
		}

		var amount float64
		for _, v := range items {
			amount += v.CalculateAmount()
			cart_suject.AddItem(v)
		}

		cart_suject.CalculateAmount()
		assert.Equal(t, amount, cart_suject.Amount)
	})
}
