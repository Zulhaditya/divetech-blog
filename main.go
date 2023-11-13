package main

import (
	"net/http"

	"github.com/Zulhaditya/divetech-blog/initdb"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialized new gin router
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		// load .env
		initdb.LoadEnv()
		// new database connection
		initdb.NewDatabase()
		ctx.JSON(http.StatusOK, gin.H{"data": "hello world!"})
	})
	router.Run(":8000")
}
