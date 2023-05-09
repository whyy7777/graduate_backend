package db

import (
	"fmt"
	"music_web/common"
	"strconv"
	"time"
)

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
		sqlStr = `SELECT id, song_name, singer, release_date, album, time, song_id FROM songs WHERE id = '` + id + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time, &temp.SongId)
		res = append(res, temp)
	}
	return res
}
