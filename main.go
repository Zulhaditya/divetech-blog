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

	// blog
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	// user
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Post{}, &models.User{})
	router.Gin.Run(":8000")

}
