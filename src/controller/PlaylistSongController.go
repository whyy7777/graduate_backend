package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
)

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

func GetPlaylist(context *gin.Context) {
	playlistId := context.Query("playlistId")
	playlist := db.GetPlaylistDetails(playlistId)
	songs := db.GetPlaylist(playlistId)
	context.JSON(200, gin.H{
		"msg":     "query success",
		"code":    "200",
		"data":    songs,
		"details": playlist,
	})
}
