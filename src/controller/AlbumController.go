package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
)

func GetAlbums(context *gin.Context) {
	_, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	singer := context.Query("singerName")
	albums := db.GetAlbums(singer)
	context.JSON(200, gin.H{
		"msg":  "get success",
		"data": albums,
	})
}
