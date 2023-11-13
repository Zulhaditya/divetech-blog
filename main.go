package main

import (
	"github.com/Zulhaditya/divetech-blog/initdb"
	"github.com/gin-gonic/gin"
	"net/http"
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
