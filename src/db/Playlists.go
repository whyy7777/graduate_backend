package db

import (
	"fmt"
	"music_web/common"
	"strconv"
	"time"
)

func NewPlaylist(uid uint, playlistName string) {
	valid := -1
	nowTime := time.Now()
	date := strconv.Itoa(nowTime.Year()) + "-" + strconv.Itoa(int(nowTime.Month())) + "-" + strconv.Itoa(nowTime.Day())
	sqlStr := `SELECT playlistId FROM playlists WHERE userId = ` + strconv.Itoa(int(uid)) + ` && playlistName = '` + playlistName + `';`
	db.QueryRow(sqlStr).Scan(&valid)
	if valid != -1 {
		return
	}
	sqlStr = `INSERT INTO playlists(userId, playlistName, establish_date )VALUES('` + strconv.Itoa(int(uid)) + `','` + playlistName + `','` + date + `');`
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func DeletePlaylist(uid uint, playlistName string) {
	sqlStr := `DELETE FROM playlists WHERE userId = ` + strconv.Itoa(int(uid)) + ` && playlistName = '` + playlistName + `';`
	fmt.Println(sqlStr)
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
}

func GetPlaylists(userId uint) []common.Playlist {
	res := make([]common.Playlist, 0)
	sqlStr := `SELECT * FROM playlists WHERE userId = '` + strconv.Itoa(int(userId)) + `';`
	playlists, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for playlists.Next() {
		var temp common.Playlist
		playlists.Scan(&temp.PlaylistId, &temp.UserId, &temp.PlaylistName, &temp.EstablishDate)
		res = append(res, temp)
	}
	return res
}

func GetPlaylist(playlistId string) []common.Song {
	res := make([]common.Song, 0)
	sqlStr := `SELECT songId FROM playlist_songs WHERE playlistId = '` + playlistId + `';`
	songId, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songId.Next() {
		var id string
		songId.Scan(&id)
		sqlStr = `SELECT * FROM songs WHERE id = '` + id + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time)
		res = append(res, temp)
	}
	return res
}

func AddToPlaylist(playlistId string, songId string) {
	var id = -1
	sqlStr := `SELECT id FROM playlist_songs WHERE playlistId = ` + playlistId + ` && songId = '` + songId + `';`
	db.QueryRow(sqlStr).Scan(&id)
	if id != -1 {
		return
	}

	sqlStr = `INSERT INTO playlist_songs(playlistId, songId)VALUES('` + playlistId + `','` + songId + `');`
	db.Exec(sqlStr)
}

func DeleteFromPlaylist(playlistId string, songId string) {
	sqlStr := `DELETE FROM playlist_songs WHERE playlistId = ` + playlistId + ` && songId = '` + songId + `';`
	db.Exec(sqlStr)
}

func GetHotPlaylists(userId string) []common.Playlist {
	data := make([]common.Playlist, 0)
	sqlStr := `SELECT playlistId FROM hot_playlists WHERE userId = '` + userId + `';`
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
func GetRecommendPlaylists(userId string) []common.Playlist {
	data := make([]common.Playlist, 0)
	sqlStr := `SELECT playlistId FROM hot_playlists WHERE userId = '` + userId + `';`
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

func GetPlaylistDetails(playlistId string) common.Playlist {
	var res common.Playlist
	sqlStr := `SELECT * FROM playlists WHERE playlistId = ` + playlistId + `;`
	db.QueryRow(sqlStr).Scan(&res.PlaylistId, &res.UserId, &res.PlaylistName, &res.EstablishDate, &res.SongCount, &res.PlayCount)
	return res
}
