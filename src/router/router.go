package router

import (
	"github.com/gin-gonic/gin"
	"music_web/db"
)

func NetInit() *gin.Engine {
	r := gin.Default()
	r.POST("/register", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		gender := context.PostForm("gender")
		ret := db.Register(username, password, gender)
		context.JSON(200, gin.H{
			"result": ret,
		})
	})
	r.POST("/login", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		ret := db.Validate(username, password)
		context.JSON(200, gin.H{
			"result": ret,
		})
	})
	return r
}
