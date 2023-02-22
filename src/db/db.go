package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

var db *sql.DB

func InitDB() (err error) {
	dsn := "root:613181hyy@tcp(127.0.0.1:3306)/music_online"
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

func Register(username, password, gender string) {
	nowTime := time.Now()
	date := strconv.Itoa(nowTime.Year()) + "-" + strconv.Itoa(int(nowTime.Month())) + "-" + strconv.Itoa(nowTime.Day())
	fmt.Printf(date)
	sqlStr := `INSERT INTO users(username, password, gender, register_time)VALUES('` + username + `','` + password + `','` + gender + `','` + date + `');`
	fmt.Println(sqlStr)
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func Validate(username, password string) bool {
	var realPassword string
	sqlStr := `SELECT password FROM users WHERE username = '` + username + `';`
	err := db.QueryRow(sqlStr).Scan(&realPassword)
	if err != nil {
		return false
	}
	fmt.Println(realPassword)
	if realPassword == password {
		return true
	}
	return false
}
