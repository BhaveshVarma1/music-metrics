package dal

import (
	"database/sql"
	"fmt"
	"music-metrics-back/model"
)

func CreateSong(tx *sql.Tx, song *model.Song) error {
	_, err := tx.Exec("INSERT INTO song (id, name, artist, album, explicit, popularity, duration, year) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		song.Id, song.Name, song.Artist, song.Album, song.Explicit, song.Popularity, song.Duration, song.Year)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveSong(tx *sql.Tx, id string) (model.Song, error) {
	rows, err := tx.Query("SELECT * FROM song WHERE id = ?;", id)
	if err != nil {
		return model.Song{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var song model.Song
	for rows.Next() {
		err = rows.Scan(&song.Id, &song.Name, &song.Artist, &song.Album, &song.Explicit, &song.Popularity, &song.Duration, &song.Year)
		if err != nil {
			return model.Song{}, err
		}
		return song, nil
	}
	return model.Song{}, nil
}

func UpdateSong(tx *sql.Tx, song *model.Song) error {
	_, err := tx.Exec("UPDATE song SET name = ?, artist = ?, album = ?, explicit = ?, popularity = ?, duration = ?, year = ? WHERE id = ?;",
		song.Name, song.Artist, song.Album, song.Explicit, song.Popularity, song.Duration, song.Year, song.Id)
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

func ClearSongs(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM song;")
	if err != nil {
		return err
	}
	return nil
}
