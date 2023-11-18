package routes

import (
	"controller"
	"initdb"
)

// route for user module
type UserRoute struct {
	Handler    initdb.GinRouter
	Controller controller.UserController
}

// initializes new instance of UserRoute
func NewUserRoute(
	controller controller.UserController,
	handler initdb.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// setup user routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.Login)
	}
}
