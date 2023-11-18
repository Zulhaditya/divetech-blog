package initdb

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin router
type GinRouter struct {
	Gin *gin.Engine
}

// new gin router all the routes defined here
func NewGinRouter() GinRouter {
	httpRouter := gin.Default()

	httpRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Up and running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}
}
