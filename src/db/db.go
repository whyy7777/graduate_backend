package db

import (
	"database/sql"
	"fmt"
	"music_web/common"
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

func QuerySong(userId uint) []common.Song {
	res := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM user_recommend WHERE userId = '` + strconv.Itoa(int(userId)) + `';`
	songs, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songs.Next() {
		var songId int
		songs.Scan(&songId)
		sqlStr = `SELECT * FROM songs WHERE id = '` + strconv.Itoa(songId) + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate)
		res = append(res, temp)
	}
	return res
}

func InsertLike(id int, songID uint) {
	sqlStr := `INSERT INTO likes(userId, songId)VALUES('` + strconv.Itoa(id) + `','` + strconv.Itoa(int(songID)) + `');`
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteLike(id int, songID uint) {
	sqlStr := `DELETE FROM likes WHERE userId = ` + strconv.Itoa(id) + ` && songId = ` + strconv.Itoa(int(songID)) + `;`
	fmt.Println(sqlStr)
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func GetLike(userId uint) []common.Song {
	res := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM likes WHERE userId = '` + strconv.Itoa(int(userId)) + `';`
	songs, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songs.Next() {
		var songId int
		songs.Scan(&songId)
		sqlStr = `SELECT * FROM songs WHERE id = '` + strconv.Itoa(songId) + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate)
		res = append(res, temp)
	}
	return res
}
