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
	r.GET("/recommend_song", middleware.AuthMiddleware(), controller.QueryRecommend)
	r.POST("/like", middleware.AuthMiddleware(), controller.AddLike)
	r.DELETE("like", middleware.AuthMiddleware(), controller.DeleteLike)
	r.GET("like", middleware.AuthMiddleware(), controller.GetLike)
	r.POST("playlist", middleware.AuthMiddleware(), controller.NewPlaylist)
	r.GET("playlist", middleware.AuthMiddleware(), controller.GetPlaylist)
	r.DELETE("playlist", middleware.AuthMiddleware(), controller.DeletePlayList)
	r.GET("playlists", middleware.AuthMiddleware(), controller.GetPlaylists)
	return r
}
