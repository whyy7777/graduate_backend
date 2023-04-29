package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music_web/common"
	"music_web/db"
	"strconv"
)

func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)
	ret, id := db.Validate(username, password)
	switch ret {
	case 0:
		token, err := common.ReleaseToken(id)
		if err != nil {
			context.JSON(500, gin.H{
				"code": 2,
				"msg":  "internal error",
			})
			return
		} else {
			context.JSON(200, gin.H{
				"code": 0,
				"msg":  "login success",
				"data": gin.H{
					"token": token,
				},
			})
		}

	case 1:
		context.JSON(404, gin.H{
			"code": 1,
			"msg":  "user doesn't exist",
		})
	case 2:
		context.JSON(500, gin.H{
			"code": 2,
			"msg":  "internal error",
		})
	case 3:
		context.JSON(500, gin.H{
			"code": 3,
			"msg":  "incorrect password",
		})
	}
}

func Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	gender := context.PostForm("gender")
	ret := db.Register(username, password, gender)
	switch ret {
	case 0:
		context.JSON(200, gin.H{
			"code": 0,
			"msg":  "register success",
		})
	case 1:
		context.JSON(200, gin.H{
			"code": 1,
			"msg":  "username is not available",
		})
	case 2:
		context.JSON(200, gin.H{
			"code": 2,
			"msg":  "internal error",
		})
	}
}

func Info(context *gin.Context) {
	user, _ := context.Get("user")
	fmt.Println(strconv.Itoa(int(user.(uint))))
	data := db.GetUserInfo(strconv.Itoa(int(user.(uint))))
	context.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"user": data,
		},
	})
}
