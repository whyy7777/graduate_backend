package main

import (
	"github.com/gin-gonic/gin"
	"music_web/controller"
	"music_web/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	user := r.Group("/user")

	user.POST("/register", controller.Register)
	user.POST("/login", controller.Login)
	user.GET("/info", middleware.AuthMiddleware(), controller.Info)

	r.POST("/like", middleware.AuthMiddleware(), controller.AddLike)
	r.DELETE("/like", middleware.AuthMiddleware(), controller.DeleteLike)
	r.GET("/like", middleware.AuthMiddleware(), controller.GetLike)

	r.POST("/playlist", middleware.AuthMiddleware(), controller.NewPlaylist)
	r.GET("/playlist", middleware.AuthMiddleware(), controller.GetPlaylist)
	r.DELETE("/playlist", middleware.AuthMiddleware(), controller.DeletePlayList)

	r.GET("/playlists/recommend", middleware.AuthMiddleware(), controller.GetRecommendPlaylists)
	r.GET("/playlists/hot", middleware.AuthMiddleware(), controller.GetHotPlaylists)

	r.PUT("/playlist_song", middleware.AuthMiddleware(), controller.AddToPlaylist)
	r.DELETE("/playlist_song", middleware.AuthMiddleware(), controller.DeleteFromPlaylist)

	r.GET("/playlists", middleware.AuthMiddleware(), controller.GetPlaylists)
	r.GET("/recommend_song", middleware.AuthMiddleware(), controller.QueryRecommend)

	r.GET("/album", middleware.AuthMiddleware(), controller.GetAlbumSongs)
	r.GET("/albums", middleware.AuthMiddleware(), controller.GetAlbums)
	return r
}
