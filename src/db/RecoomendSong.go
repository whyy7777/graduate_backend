package db

import (
	"music_web/common"
	"strconv"
)

func RecommendSong(userId uint) []common.Song {
	res := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM recommend_songs WHERE userId = '` + strconv.Itoa(int(userId)) + `';`
	songs, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songs.Next() {
		var songId int
		songs.Scan(&songId)
		sqlStr = `SELECT id, song_name, singer, release_date, album, time FROM songs WHERE id = '` + strconv.Itoa(songId) + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time)
		res = append(res, temp)
	}
	return res
}
