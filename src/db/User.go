package db

import (
	"music_web/common"
	"strconv"
	"time"
)

// Register 0: register success, 1: already have same username, 2: internal error
func Register(username, password, gender string) int {
	var id = -1
	db.QueryRow("SELECT id FROM users WHERE username = '" + username + "'").Scan(&id)
	if id != -1 {
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
func Validate(username, password string) (int, uint) {
	var id = -1
	db.QueryRow("SELECT id FROM users WHERE username = '" + username + "'").Scan(&id)
	if id == -1 {
		return 1, 0
	}
	var realPassword string
	sqlStr := `SELECT password FROM users WHERE username = '` + username + `';`
	err := db.QueryRow(sqlStr).Scan(&realPassword)
	if err != nil {
		return 2, 0
	}
	if realPassword == password {
		return 0, uint(id)
	}
	return 3, 0
}

func GetUserInfo(userId string) common.User {
	sqlStr := `SELECT * FROM users WHERE id = ` + userId + `;`
	var res common.User
	db.QueryRow(sqlStr).Scan(&res.Id, &res.Username, &res.Password, &res.Gender, &res.RegisterTime, &res.CountFollow, &res.CountFollowed, &res.Description, &res.CountMoment, &res.CountCreatePlaylist, &res.CountLikePlaylist, &res.Level)
	res.Password = " "
	return res
}
