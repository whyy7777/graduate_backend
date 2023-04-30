package db

import "music_web/common"

func GetRecommendPlaylists(userId string) []common.Playlist {
	data := make([]common.Playlist, 0)
	sqlStr := `SELECT playlistId FROM recommend_playlists WHERE userId = '` + userId + `';`
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
		playlist.Scan(&temp.PlaylistId, &temp.UserId, &temp.PlaylistName, &temp.EstablishDate, &temp.SongCount, &temp.PlayCount)
		data = append(data, temp)
	}
	return data
}
