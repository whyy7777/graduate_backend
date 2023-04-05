package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"music_web/db"
)

var p = flag.String("p", "", "password of db")

func main() {
	flag.Parse()
	password := *p
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
