package cart

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(g *gin.Engine) {
	g.POST("/carts", func(ctx *gin.Context) {
		var body struct {
			Currency string      `json:"currency"`
			User     interface{} `json:"user"`
		}
		if err := ctx.BindJSON(&body); err != nil {
			log.Println("error binding body: " + err.Error())
			ctx.JSON(400, gin.H{
				"status": "Cart not created",
				"error":  err,
			})
			return
		}
		cart := CreateCart(body.Currency)
		if err := cart.Persist(); err != nil {
			log.Println("error during persist cart: " + err.Error())
			ctx.JSON(400, gin.H{
				"status": "Cart not created",
				"error":  err,
			})
			return
		}
		ctx.JSON(201, gin.H{
			"status": "cart created",
			"data":   cart,
		})
	})

	g.GET("/carts/:cart_id", func(ctx *gin.Context) {
		cart := FindCartById(ctx.Param("cart_id"))
		if cart != nil {
			cart.CalculateAmount()
		}
		ctx.JSON(200, gin.H{
			"cart": cart,
		})
	})

	g.POST("/carts/:cart_id/items", func(ctx *gin.Context) {
		var item Item
		if err := ctx.BindJSON(&item); err != nil {
			log.Println(err)
			ctx.JSON(400, gin.H{
				"status": "error adding item to cart",
				"error":  err.Error(),
			})
			return
		}
		cart := FindCartById(ctx.Param("cart_id"))
		if cart == nil {
			log.Println(ctx.Param("cart_id"))
			ctx.JSON(400, gin.H{
				"status": "error adding item to cart",
				"error":  "Cart not found",
			})
			return
		}
		if err := cart.AddItemAndSave(item); err != nil {
			log.Println(err)
			ctx.JSON(400, gin.H{
				"status": "error adding item to cart",
				"error":  err.Error(),
			})
			return
		}
		ctx.JSON(201, gin.H{
			"status": "item added to cart",
			"data":   cart,
		})
	})
}
