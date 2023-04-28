package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
	"strconv"
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
