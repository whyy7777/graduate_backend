package db

func AddToPlaylist(playlistId string, songId string) {
	var id = -1
	sqlStr := `SELECT id FROM playlist_songs WHERE playlistId = ` + playlistId + ` && songId = '` + songId + `';`
	err := db.QueryRow(sqlStr).Scan(&id)
	if err != nil {
		return
	}
	if id != -1 {
		return
	}
	sqlStr = `INSERT INTO playlist_songs(playlistId, songId)VALUES('` + playlistId + `','` + songId + `');`
	_, err = db.Exec(sqlStr)
	if err != nil {
		return
	}
}

func DeleteFromPlaylist(playlistId string, songId string) {
	sqlStr := `DELETE FROM playlist_songs WHERE playlistId = ` + playlistId + ` && songId = '` + songId + `';`
	_, err := db.Exec(sqlStr)
	if err != nil {
		return
	}
}
