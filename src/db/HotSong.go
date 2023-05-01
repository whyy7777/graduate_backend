package db

import (
	"music_web/common"
	"strconv"
)

func GetHotSongs(userId string) []common.Song {
	res := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM hot_songs WHERE userId = '` + userId + `';`
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
