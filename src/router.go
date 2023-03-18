package main

import (
	"github.com/gin-gonic/gin"
	"music_web/controller"
	"music_web/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
