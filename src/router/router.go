package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music_web/common"
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
		token, err := common.ReleaseToken(10)
		if err != nil {
			context.JSON(500, gin.H{
				"msg": "internal error",
			})
			fmt.Printf("%v\n", err)
			return
		} else {
			context.JSON(200, gin.H{
				"result": ret,
				"data": gin.H{
					"token": token,
				},
			})
		}

	})
	return r
}
