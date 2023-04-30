package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/common"
	"music_web/db"
	"strconv"
)

func GetHotSong(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	data := make([]common.Song, 0)
	data = db.GetHotSongs(strconv.Itoa(int(id.(uint))))
	context.JSON(200, gin.H{
		"msg":  " query success",
		"data": data,
	})
}
