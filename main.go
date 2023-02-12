package main

import (
	"gin-service/middleware"
	"gin-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ping test
	r.GET("/ping", service.GetPong)

	// middleware
	r.Use(middleware.LoginMiddleware())

	// group
	route := r.Group("/user")
	{
		route.POST("/getuser", service.GetUser)
		route.POST("/showuser", service.ShowUser)
	}

	r.Run()
}
