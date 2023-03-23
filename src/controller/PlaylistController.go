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

func GetPlaylist(context *gin.Context) {
	playlistId := context.Query("playlistId")
	songs := db.GetPlaylist(playlistId)
	context.JSON(200, gin.H{
		"msg":  "query success",
		"code": "200",
		"data": songs,
	})
}

func AddToPlaylist(context *gin.Context) {
	_, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	songId := context.Query("songId")
	playlistId := context.Query("playlistId")
	db.AddToPlaylist(playlistId, songId)
	context.JSON(200, gin.H{
		"msg":  "execute success",
		"code": "200",
	})
}

func GetPlaylists(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	playlists := db.GetPlaylists(id.(uint))
	context.JSON(200, gin.H{
		"msg":  "get success",
		"data": playlists,
	})

}

func DeleteFromPlaylist(context *gin.Context) {
	_, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	songId := context.Query("songId")
	playlistId := context.Query("playlistId")
	db.DeleteFromPlaylist(playlistId, songId)
	context.JSON(200, gin.H{
		"msg":  "execute success",
		"code": "200",
	})
}
