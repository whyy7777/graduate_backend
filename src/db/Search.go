package db

import (
	"music_web/common"
)

func SearchSong(songName string) []common.Song {
	res := make([]common.Song, 0)

	sqlStr := `SELECT id FROM songs WHERE song_name LIKE '%` + songName + `%';`
	songs, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for songs.Next() {
		var songId string
		err = songs.Scan(&songId)
		if err != nil {
			return nil
		}
		sqlStr = `SELECT id, song_name, singer, release_date, album, time, song_id FROM songs WHERE id = '` + songId + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Song
		err = song.Scan(&temp.Id, &temp.SongName, &temp.Singer, &temp.ReleaseDate, &temp.Album, &temp.Time, &temp.SongId)
		if err != nil {
			return nil
		}
		res = append(res, temp)
	}
	return res
}

func SearchSinger(singerName string) []common.Singer {
	res := make([]common.Singer, 0)
	sqlStr := `SELECT singerId FROM singer WHERE singerName LIKE '%` + singerName + `%';`
	singers, err := db.Query(sqlStr)
	if err != nil {
		return nil
	}
	for singers.Next() {
		var singerId string
		err = singers.Scan(&singerId)
		if err != nil {
			return nil
		}
		sqlStr = `SELECT singerId, singerName, gender, bornDate FROM songs WHERE id = '` + singerId + `';`
		song := db.QueryRow(sqlStr)
		var temp common.Singer
		err = song.Scan(&temp.SingerId, &temp.SingerName, &temp.Gender, &temp.BornDate)
		if err != nil {
			return nil
		}
		res = append(res, temp)
	}
	return res
}
