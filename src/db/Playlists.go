package db

import (
	"music_web/common"
	"strconv"
)

func GetCreatePlaylists(userId uint) []common.Playlist {
	res := make([]common.Playlist, 0)
	sqlStr := `SELECT * FROM playlists WHERE userId = '` + strconv.Itoa(int(userId)) + `';`
	playlists, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for playlists.Next() {
		var temp common.Playlist
		err = playlists.Scan(&temp.PlaylistId, &temp.UserId, &temp.PlaylistName, &temp.EstablishDate, &temp.SongCount, &temp.PlayCount)
		if err != nil {
			return nil
		}
		res = append(res, temp)
	}
	return res
}

func GetLikePlaylists(userId string) []common.Playlist {
	data := make([]common.Playlist, 0)
	sqlStr := `SELECT playlistId FROM like_playlists WHERE userId = '` + userId + `';`
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

func GetPlaylistDetails(playlistId string) common.Playlist {
	var res common.Playlist
	sqlStr := `SELECT * FROM playlists WHERE playlistId = ` + playlistId + `;`
	err := db.QueryRow(sqlStr).Scan(&res.PlaylistId, &res.UserId, &res.PlaylistName, &res.EstablishDate, &res.SongCount, &res.PlayCount)
	if err != nil {
		return res
	}
	return res
}
