package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"music_web/db"
	"os"
)

func main() {
	password := os.Args[1]
	r := CollectRoute(gin.Default())
	err := db.InitDB(password)
	if err != nil {
		fmt.Printf("init DB failed,err%v\n", err)
	}
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
