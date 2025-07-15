package http

import (
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController *UserHandler
}

func (R *RouteConfig) Setup() {
	R.PublicRoutes()
	R.AuthRoutes()
}

func (R *RouteConfig) PublicRoutes() {
	R.App.POST("/signup", R.UserController.Signup)
	R.App.GET("/users", R.UserController.GetUsers)

}

func (R *RouteConfig) AuthRoutes() {

}
