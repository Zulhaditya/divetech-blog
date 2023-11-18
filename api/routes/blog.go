package routes

import (
	"controller"
	"initdb"
)

// routes for question module
type PostRoute struct {
	Controller controller.PostController
	Handler    initdb.GinRouter
}

// initialize new choice routes
func NewPostRoute(
	controller controller.PostController,
	handler initdb.GinRouter,
) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// setup new choice routes
func (p PostRoute) Setup() {
	// router group
	post := p.Handler.Gin.Group("/posts")
	{
		post.GET("/", p.Controller.GetPosts)
		post.POST("/", p.Controller.AddPost)
		post.GET("/:id", p.Controller.GetPost)
		post.DELETE("/:id", p.Controller.DeletePost)
		post.PUT("/:id", p.Controller.UpdatedPost)
	}
}
