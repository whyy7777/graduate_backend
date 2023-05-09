package db

import (
	"fmt"
	"music_web/common"
	"strconv"
)

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
		sqlStr = `SELECT id, song_name, singer, release_date, album, time, song_id FROM songs WHERE id = '` + strconv.Itoa(songId) + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time, &temp.SongId)
		res = append(res, temp)
	}
	return res
}
