package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
)

func NewPlaylist(context *gin.Context) {
	playlistName := context.PostForm("playlistName")
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	db.NewPlaylist(id.(uint), playlistName)
	context.JSON(200, gin.H{
		"msg":          "add success",
		"code":         200,
		"playlistName": playlistName,
	})
}

func DeletePlayList(context *gin.Context) {
	playlistName := context.PostForm("playlistName")
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	db.DeletePlaylist(id.(uint), playlistName)
	context.JSON(200, gin.H{
		"msg":          "delete success",
		"playlistName": playlistName,
	})
}
