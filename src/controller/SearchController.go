package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/common"
	"music_web/db"
)

func SearchSong(context *gin.Context) {
	_, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	songName := context.Query("keyword")
	data := make([]common.Song, 0)
	data = db.SearchSong(songName)
	context.JSON(200, gin.H{
		"msg":  " query success",
		"data": data,
	})
}

func SearchSinger(context *gin.Context) {
	_, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	singerName := context.Query("keyword")
	data := make([]common.Singer, 0)
	data = db.SearchSinger(singerName)
	context.JSON(200, gin.H{
		"msg":  " query success",
		"data": data,
	})
}
