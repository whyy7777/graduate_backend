package db

import (
	"fmt"
	"music_web/common"
)

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
	sqlStr := `SELECT songId FROM album_songs WHERE albumId = '` + albumId + `';`
	songIds, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songIds.Next() {
		var temp common.Song
		var id string
		songIds.Scan(&id)
		fmt.Println(id)
		sqlStr = `SELECT * FROM songs WHERE id = ` + id + `;`
		fmt.Println(sqlStr)
		song := db.QueryRow(sqlStr)
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time)
		songs = append(songs, temp)
	}
	return songs
}
