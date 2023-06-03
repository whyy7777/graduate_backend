package db

import (
	"music_web/common"
	"strconv"
)

func GetHotSongs(userId string) []common.Song {
	res := make([]common.Song, 0)
	userId = "9"
	sqlStr := `SELECT songId FROM hot_songs WHERE userId = '` + userId + `';`
	songs, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songs.Next() {
		var songId int
		err = songs.Scan(&songId)
		if err != nil {
			return nil
		}
		sqlStr = `SELECT id, song_name, singer, release_date, album, time, song_id FROM songs WHERE id = '` + strconv.Itoa(songId) + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		err = song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time, &temp.SongId)
		if err != nil {
			return nil
		}
		res = append(res, temp)
	}
	return res
}
