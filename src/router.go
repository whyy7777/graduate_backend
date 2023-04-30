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
	r.DELETE("/playlist", middleware.AuthMiddleware(), controller.DeletePlayList)
	r.GET("/playlist", middleware.AuthMiddleware(), controller.GetPlaylist)

	r.GET("/playlists", middleware.AuthMiddleware(), controller.GetCreatePlaylists)
	r.GET("/playlists/like", middleware.AuthMiddleware(), controller.GetLikePlaylists)

	r.PUT("/playlist/song", middleware.AuthMiddleware(), controller.AddToPlaylist)
	r.DELETE("/playlist/song", middleware.AuthMiddleware(), controller.DeleteFromPlaylist)

	r.GET("/songs/recommend", middleware.AuthMiddleware(), controller.GetRecommendSong)
	r.GET("/songs/hot", middleware.AuthMiddleware(), controller.GetHotSong)

	r.GET("/album", middleware.AuthMiddleware(), controller.GetAlbumSongs)

	r.GET("/albums", middleware.AuthMiddleware(), controller.GetAlbums)
	r.GET("/albums/hot", middleware.AuthMiddleware(), controller.GetHotAlbums)

	r.GET("/playlists/recommend", middleware.AuthMiddleware(), controller.GetRecommendPlaylists)
	r.GET("/playlists/hot", middleware.AuthMiddleware(), controller.GetHotPlaylists)
	return r
}
