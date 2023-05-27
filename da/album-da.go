package da

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateAlbum(tx *sql.Tx, album *model.AlbumBean) error {
	_, err := tx.Exec("INSERT INTO album (`id`, `name`, `artist`, `genre`, `totalTracks`, `year`, `image`, `popularity`, `artistID`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);",
		album.Id, album.Name, album.Artist, album.Genre, album.TotalTracks, album.Year, album.Image, album.Popularity, album.ArtistId)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAlbum(tx *sql.Tx, id string) (model.AlbumBean, error) {
	rows, err := tx.Query("SELECT * FROM album WHERE id = ?;", id)
	if err != nil {
		return model.AlbumBean{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var album model.AlbumBean
	for rows.Next() {
		err = rows.Scan(&album.Id, &album.Name, &album.Artist, &album.Genre, &album.TotalTracks, &album.Year, &album.Image, &album.Popularity, &album.ArtistId)
		if err != nil {
			return model.AlbumBean{}, err
		}
		return album, nil
	}
	return model.AlbumBean{}, nil
}

func UpdateAlbum(tx *sql.Tx, album *model.AlbumBean) error {
	_, err := tx.Exec("UPDATE album SET name = ?, artist = ?, genre = ?, totalTracks = ?, year = ?, image = ?, popularity = ?, artistID = ? WHERE id = ?;",
		album.Name, album.Artist, album.Genre, album.TotalTracks, album.Year, album.Image, album.Popularity, album.ArtistId, album.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAlbum(tx *sql.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM album WHERE id = ?;", id)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAllAlbums(tx *sql.Tx) ([]model.AlbumBean, error) {
	rows, err := tx.Query("SELECT * FROM album;")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var albums []model.AlbumBean
	for rows.Next() {
		var album model.AlbumBean
		err = rows.Scan(&album.Id, &album.Name, &album.Artist, &album.Genre, &album.TotalTracks, &album.Year, &album.Image, &album.Popularity, &album.ArtistId)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func ClearAlbums(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM album;")
	if err != nil {
		return err
	}
	return nil
}
