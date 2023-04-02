package dal

import (
	"database/sql"
	"fmt"
	"math"
	"music-metrics/model"
)

func GetAverageYear(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT AVG(year) FROM listen INNER JOIN song ON listen.songID = song.id INNER JOIN album ON song.album = album.id WHERE username = ?;")
	if err != nil {
		return 0, err
	}

	var result float32
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return int(math.Round(float64(result))), nil
}

func GetSongCounts(tx *sql.Tx, username string) ([]model.SongCount, error) {
	stmt, err := tx.Prepare("SELECT song.name, song.artist, COUNT(*) FROM song INNER JOIN listen ON song.id = listen.songID WHERE username = ? GROUP BY song.id ORDER BY COUNT(*) DESC;")
	if err != nil {
		return nil, err
	}

	var results []model.SongCount
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)

	for rows.Next() {
		var song string
		var artist string
		var count int
		err = rows.Scan(&song, &artist, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.SongCount{Song: song, Artist: artist, Count: count})
	}

	return results, nil
}

func GetTopAlbums(tx *sql.Tx, username string) ([]model.TopAlbum, error) {
	stmt, err := tx.Prepare("SELECT album.name, album.artist, album.image, COUNT(*) FROM album INNER JOIN song ON album.id = song.album INNER JOIN listen ON song.id = listen.songID WHERE username = ? GROUP BY album.id ORDER BY COUNT(*) DESC;")
	if err != nil {
		return nil, err
	}

	var results []model.TopAlbum
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)

	for rows.Next() {
		var album string
		var artist string
		var image string
		var count int
		err = rows.Scan(&album, &artist, &image, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.TopAlbum{Album: album, Artist: artist, Image: image, Count: count})
	}

	return results, nil
}
