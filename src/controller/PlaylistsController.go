package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
	"strconv"
)

func GetRecommendPlaylists(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	data := db.GetRecommendPlaylists(strconv.Itoa(int(id.(uint))))
	context.JSON(200, gin.H{
		"msg":  "query success",
		"data": data,
	})
}

func GetHotPlaylists(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	data := db.GetHotPlaylists(strconv.Itoa(int(id.(uint))))
	context.JSON(200, gin.H{
		"msg":  "query success",
		"data": data,
	})
}

func GetLikePlaylists(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	data := db.GetLikePlaylists(strconv.Itoa(int(id.(uint))))
	context.JSON(200, gin.H{
		"msg":  "query success",
		"data": data,
	})
}

func GetCreatePlaylists(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	playlists := db.GetCreatePlaylists(id.(uint))
	context.JSON(200, gin.H{
		"msg":  "get success",
		"data": playlists,
	})
}
