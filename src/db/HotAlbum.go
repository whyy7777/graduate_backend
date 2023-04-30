package db

import "music_web/common"

func GetHotAlbums(userId string) []common.Album {
	res := make([]common.Album, 0)
	sqlStr := `SELECT albumId FROM hot_albums WHERE userId = ` + userId + `;`
	albums, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for albums.Next() {
		var albumId string
		albums.Scan(&albumId)
		var temp common.Album
		sqlStr = `SELECT * FROM albums WHERE albumId = ` + albumId + `;`
		db.QueryRow(sqlStr).Scan(&temp.AlbumId, &temp.AlbumName, &temp.Singer, &temp.ReleaseData)
	}
	return res
}
