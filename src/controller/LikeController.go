package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
	"strconv"
)

func AddLike(context *gin.Context) {
	songID := context.PostForm("songId")
	songId, err := strconv.Atoi(songID)
	if err != nil {
		context.JSON(200, gin.H{
			"msg":  "internal error",
			"code": 404,
		})
		return
	}
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	db.InsertLike(int(id.(uint)), uint(songId))
	context.JSON(200, gin.H{
		"msg":  "add success",
		"code": 200,
	})
}

func DeleteLike(context *gin.Context) {
	songID := context.PostForm("songID")
	songId, err := strconv.Atoi(songID)
	if err != nil {
		context.JSON(200, gin.H{
			"msg":  "internal error",
			"code": 404,
		})
		return
	}
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	db.DeleteLike(int(id.(uint)), uint(songId))
	context.JSON(200, gin.H{
		"msg":  "delete success",
		"code": 200,
	})
}

func GetLike(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	data := db.GetLike(id.(uint))
	context.JSON(200, gin.H{
		"msg":  "query success",
		"data": data,
	})

}
