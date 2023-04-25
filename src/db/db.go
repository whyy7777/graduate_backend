package db

import (
	"database/sql"
	"fmt"
	"music_web/common"
	"strconv"
	"time"
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

func NewPlaylist(uid uint, playlistName string) {
	valid := -1
	nowTime := time.Now()
	date := strconv.Itoa(nowTime.Year()) + "-" + strconv.Itoa(int(nowTime.Month())) + "-" + strconv.Itoa(nowTime.Day())
	sqlStr := `SELECT playlistId FROM playlists WHERE userId = ` + strconv.Itoa(int(uid)) + ` && playlistName = '` + playlistName + `';`
	db.QueryRow(sqlStr).Scan(&valid)
	if valid != -1 {
		return
	}
	sqlStr = `INSERT INTO playlists(userId, playlistName, establish_date )VALUES('` + strconv.Itoa(int(uid)) + `','` + playlistName + `','` + date + `');`
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func DeletePlaylist(uid uint, playlistName string) {
	sqlStr := `DELETE FROM playlists WHERE userId = ` + strconv.Itoa(int(uid)) + ` && playlistName = '` + playlistName + `';`
	fmt.Println(sqlStr)
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func GetPlaylists(userId uint) []common.Playlist {
	res := make([]common.Playlist, 0)
	sqlStr := `SELECT * FROM playlists WHERE userId = '` + strconv.Itoa(int(userId)) + `';`
	playlists, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for playlists.Next() {
		var temp common.Playlist
		playlists.Scan(&temp.PlaylistId, &temp.UserId, &temp.PlaylistName, &temp.EstablishDate)
		res = append(res, temp)
	}
	return res
}

func GetPlaylist(playlistId string) []common.Song {
	res := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM playlist_songs WHERE playlistId = '` + playlistId + `';`
	songId, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songId.Next() {
		var id string
		songId.Scan(&id)
		sqlStr = `SELECT * FROM songs WHERE id = '` + id + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate)
		res = append(res, temp)
	}
	return res
}

func AddToPlaylist(playlistId string, songId string) {
	var id = -1
	sqlStr := `SELECT id FROM playlist_songs WHERE playlistId = ` + playlistId + ` && songId = '` + songId + `';`
	db.QueryRow(sqlStr).Scan(&id)
	if id != -1 {
		return
	}

	sqlStr = `INSERT INTO playlist_songs(playlistId, songId)VALUES('` + playlistId + `','` + songId + `');`
	db.Exec(sqlStr)
}

func DeleteFromPlaylist(playlistId string, songId string) {
	sqlStr := `DELETE FROM playlist_songs WHERE playlistId = ` + playlistId + ` && songId = '` + songId + `';`
	db.Exec(sqlStr)
}

func GetHotPlaylists(userId string) []common.Playlist {
	data := make([]common.Playlist, 0)
	sqlStr := `SELECT playlistId FROM hot_playlists WHERE userId = '` + userId + `';`
	playlistId, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for playlistId.Next() {
		var id string
		playlistId.Scan(&id)
		sqlStr = `SELECT * FROM playlists WHERE playlistId = '` + id + `';`
		var temp common.Playlist
		playlist := db.QueryRow(sqlStr)
		playlist.Scan(&temp.PlaylistId, &temp.UserId, &temp.PlaylistName, &temp.EstablishDate)
		data = append(data, temp)
	}
	return data
}

func GetAlbums(singer string) []common.Album {
	data := make([]common.Album, 0)
	sqlStr := `SELECT * FROM albums WHERE singer = '` + singer + `';`
	albums, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for albums.Next() {
		var temp common.Album
		albums.Scan(&temp.AlbumId, &temp.AlbumName, &temp.Singer, &temp.ReleaseData)
		data = append(data, temp)
	}
	return data
}

func GetAlbumSongs(albumId string) []common.Song {
	songs := make([]common.Song, 0)

	return songs
}
