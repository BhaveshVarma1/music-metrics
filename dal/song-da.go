package dal

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateSong(tx *sql.Tx, song *model.SongBean) error {

	_, err := tx.Exec("INSERT INTO song (`id`, `name`, `artist`, `album`, `explicit`, `popularity`, `duration`) VALUES (?, ?, ?, ?, ?, ?, ?);",
		song.Id, song.Name, song.Artist, song.Album, song.Explicit, song.Popularity, song.Duration)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveSong(tx *sql.Tx, id string) (model.SongBean, error) {
	rows, err := tx.Query("SELECT * FROM song WHERE id = ?;", id)
	if err != nil {
		return model.SongBean{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var song model.SongBean
	for rows.Next() {
		err = rows.Scan(&song.Id, &song.Name, &song.Artist, &song.Album, &song.Explicit, &song.Popularity, &song.Duration)
		if err != nil {
			return model.SongBean{}, err
		}
		return song, nil
	}
	return model.SongBean{}, nil
}

func UpdateSong(tx *sql.Tx, song *model.SongBean) error {
	_, err := tx.Exec("UPDATE song SET name = ?, artist = ?, album = ?, explicit = ?, popularity = ?, duration = ? WHERE id = ?;",
		song.Name, song.Artist, song.Album, song.Explicit, song.Popularity, song.Duration, song.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSong(tx *sql.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM song WHERE id = ?;", id)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAllSongs(tx *sql.Tx) ([]model.SongBean, error) {
	rows, err := tx.Query("SELECT * FROM song;")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var songs []model.SongBean
	for rows.Next() {
		var song model.SongBean
		err = rows.Scan(&song.Id, &song.Name, &song.Artist, &song.Album, &song.Explicit, &song.Popularity, &song.Duration)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func ClearSongs(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM song;")
	if err != nil {
		return err
	}
	return nil
}
