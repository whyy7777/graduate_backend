package db

import (
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
		err = albums.Scan(&temp.AlbumId, &temp.AlbumName, &temp.Singer, &temp.ReleaseData)
		if err != nil {
			return nil
		}
		data = append(data, temp)
	}
	return data
}
