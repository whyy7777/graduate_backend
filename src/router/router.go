package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music_web/db"
	"net/http"
)

func NetInit() *gin.Engine {
	r := gin.Default()
	r.POST("/register", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		gender := context.PostForm("gender")
		db.Register(username, password, gender)
	})
	r.POST("/login", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		res := db.Validate(username, password)
		if res {
			context.String(http.StatusOK, fmt.Sprint("right password"))
		}
	})
	return r
}
