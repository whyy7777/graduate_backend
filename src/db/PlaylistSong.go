package db

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
