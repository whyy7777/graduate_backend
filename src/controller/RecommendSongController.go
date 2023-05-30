package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music_web/common"
	"music_web/db"
	"os/exec"
	"strconv"
)

func GetRecommendSong(context *gin.Context) {
	id, ok := context.Get("user")
	if !ok {
		context.JSON(200, gin.H{
			"msg":  "login first",
			"code": 404,
		})
		return
	}
	data := make([]common.Song, 0)
	userId := strconv.Itoa(int(id.(uint)))
	cmd := exec.Command("python", "predictUser.py", userId)
	fmt.Println(cmd)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	data = db.RecommendSong(id.(uint))
	context.JSON(200, gin.H{
		"msg":  " query success",
		"data": data,
	})
}
