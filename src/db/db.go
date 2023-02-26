package db

import (
	"database/sql"
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

// Register 0: register success, 1: already have same username, 2: internal error
func Register(username, password, gender string) int {
	var id = -1
	db.QueryRow("SELECT id FROM users WHERE username = '" + username + "'").Scan(&id)
	if id == -1 {
		return 1
	}
	nowTime := time.Now()
	date := strconv.Itoa(nowTime.Year()) + "-" + strconv.Itoa(int(nowTime.Month())) + "-" + strconv.Itoa(nowTime.Day())
	sqlStr := `INSERT INTO users(username, password, gender, register_time)VALUES('` + username + `','` + password + `','` + gender + `','` + date + `');`
	_, err := db.Exec(sqlStr)
	if err != nil {
		return 2
	} else {
		return 0
	}
}

// Validate 0: log in success, 1: user doesn't exist, 2: internal error, 3: incorrect password
func Validate(username, password string) int {
	var id = -1
	db.QueryRow("SELECT id FROM users WHERE username = '" + username + "'").Scan(&id)
	if id == -1 {
		return 1
	}
	var realPassword string
	sqlStr := `SELECT password FROM users WHERE username = '` + username + `';`
	err := db.QueryRow(sqlStr).Scan(&realPassword)
	if err != nil {
		return 2
	}
	if realPassword == password {
		return 0
	}
	return 3
}
