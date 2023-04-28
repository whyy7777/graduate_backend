package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitDB(password string) (err error) {
	fmt.Println(password)
	dsn := "root:" + password + "@tcp(1.117.65.130:3306)/music_online"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxIdleConns(10)
	return
}
