package src

import (
	"github.com/EzequielBPullolil/cart_service/src/cart"
	"github.com/gin-gonic/gin"
)

func CreateApp() *gin.Engine {
	g := gin.Default()
	cart.HandleRoutes(g)
	return g
}
