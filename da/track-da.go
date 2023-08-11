package da

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateTrack(tx *sql.Tx, track *model.TrackBean) error {
	_, err := tx.Exec("INSERT INTO track (`id`, `name`, `artist`, `album`, `explicit`, `popularity`, `duration`, `artistID`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		track.Id, track.Name, track.Artist, track.Album, track.Explicit, track.Popularity, track.Duration, track.ArtistId)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveTrack(tx *sql.Tx, id string) (model.TrackBean, error) {
	rows, err := tx.Query("SELECT * FROM track WHERE id = ?;", id)
	if err != nil {
		return model.TrackBean{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var track model.TrackBean
	for rows.Next() {
		err = rows.Scan(&track.Id, &track.Name, &track.Artist, &track.Album, &track.Explicit, &track.Popularity, &track.Duration, &track.ArtistId)
		if err != nil {
			return model.TrackBean{}, err
		}
		return track, nil
	}
	return model.TrackBean{}, nil
}

func UpdateTrack(tx *sql.Tx, track *model.TrackBean) error {
	_, err := tx.Exec("UPDATE track SET name = ?, artist = ?, album = ?, explicit = ?, popularity = ?, duration = ?, artistID = ? WHERE id = ?;",
		track.Name, track.Artist, track.Album, track.Explicit, track.Popularity, track.Duration, track.ArtistId, track.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTrack(tx *sql.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM track WHERE id = ?;", id)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAllTracks(tx *sql.Tx) ([]model.TrackBean, error) {
	rows, err := tx.Query("SELECT * FROM track;")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var tracks []model.TrackBean
	for rows.Next() {
		var track model.TrackBean
		err = rows.Scan(&track.Id, &track.Name, &track.Artist, &track.Album, &track.Explicit, &track.Popularity, &track.Duration, &track.ArtistId)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func ClearTracks(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM track;")
	if err != nil {
		return err
	}
	return nil
}
