package db

import (
	"fmt"
	"music_web/common"
)

func GetAlbumSongs(albumId string) []common.Song {
	songs := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM album_songs WHERE albumId = '` + albumId + `';`
	songIds, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songIds.Next() {
		var temp common.Song
		var id string

		err = songIds.Scan(&id)
		if err != nil {
			return nil
		}
		fmt.Println(id)
		sqlStr = `SELECT id, song_name, singer, release_date, album, time, song_id FROM songs WHERE id = ` + id + `;`
		fmt.Println(sqlStr)
		song := db.QueryRow(sqlStr)
		err = song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time, &temp.SongId)
		if err != nil {
			return nil
		}
		songs = append(songs, temp)
	}
	return songs
}
