package main

import (
	"apiroutes"
	"controller"
	"initdb"
	"models"
	"repository"
	"service"
)

func init() {
	initdb.LoadEnv()
}

func main() {
	router := initdb.NewGinRouter()
	db := initdb.NewDatabase()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	db.DB.AutoMigrate(&models.Post{})
	router.Gin.Run(":8000")

}
