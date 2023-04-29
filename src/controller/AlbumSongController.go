package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
)

func GetAlbumSongs(context *gin.Context) {
	_, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	albumId := context.Query("albumId")
	songs := db.GetAlbumSongs(albumId)
	context.JSON(200, gin.H{
		"msg":  "get success",
		"data": songs,
	})
}
