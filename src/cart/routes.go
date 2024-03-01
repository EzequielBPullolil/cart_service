package cart

import (
	"github.com/gin-gonic/gin"
)

func HandleRoutes(g *gin.Engine) {
	g.POST("/carts", func(ctx *gin.Context) {
		cart := CreateCart()

		if err := cart.Persist(); err != nil {
			ctx.JSON(400, gin.H{
				"status": "Cart not created",
				"error":  err,
			})
		}
		ctx.JSON(201, gin.H{
			"status": "cart created",
			"data":   cart,
		})
	})

	g.GET("/carts/:cart_id", func(ctx *gin.Context) {
		cart := FindCartById(ctx.Param("cart_id"))
		ctx.JSON(200, gin.H{
			"cart": cart,
		})
	})
		})
	})
}
