package db

import "music_web/common"

func GetHotPlaylists(userId string) []common.Playlist {
	data := make([]common.Playlist, 0)
	sqlStr := `SELECT playlistId FROM hot_playlists WHERE userId = '` + userId + `';`
	playlistId, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for playlistId.Next() {
		var id string
		err = playlistId.Scan(&id)
		if err != nil {
			return nil
		}
		sqlStr = `SELECT * FROM playlists WHERE playlistId = '` + id + `';`
		var temp common.Playlist
		playlist := db.QueryRow(sqlStr)
		err = playlist.Scan(&temp.PlaylistId, &temp.UserId, &temp.PlaylistName, &temp.EstablishDate, &temp.SongCount, &temp.PlayCount)
		if err != nil {
			return nil
		}
		data = append(data, temp)
	}
	return data
}
