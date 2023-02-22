package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"music_web/db"
	"music_web/router"
)

func main() {
	r := router.NetInit()
	err := db.InitDB()
	if err != nil {
		fmt.Printf("init DB failed,err%v\n", err)
	}
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
