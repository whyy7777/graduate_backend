package controller

import (
	"github.com/gin-gonic/gin"
	"music_web/common"
	"music_web/db"
)

func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	ret, id := db.Validate(username, password)
	switch ret {
	case 0:
		{
			token, err := common.ReleaseToken(id)
			if err != nil {
				context.JSON(500, gin.H{
					"msg": "internal error",
				})
				return
			} else {
				context.JSON(200, gin.H{
					"msg": "login success",
					"data": gin.H{
						"token": token,
					},
				})
			}
		}
	case 1:
		context.JSON(404, gin.H{
			"msg": "user doesn't exist",
		})
	case 2:
		context.JSON(500, gin.H{
			"msg": "internal error",
		})
	case 3:
		context.JSON(500, gin.H{
			"msg": "incorrect password",
		})
	}
}

func Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	gender := context.PostForm("gender")
	ret := db.Register(username, password, gender)

	context.JSON(200, gin.H{
		"result": ret,
	})
}

func Info(context *gin.Context) {
	user, _ := context.Get("user")
	context.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"user": user,
		},
	})
}
