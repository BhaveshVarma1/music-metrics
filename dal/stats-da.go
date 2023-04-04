package dal

import (
	"database/sql"
	"fmt"
	"math"
	"music-metrics/model"
)

func GetAverageYear(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT AVG(year) FROM listen l JOIN song s ON l.songID = s.id JOIN album a ON s.album = a.id WHERE username = ?;")
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
	stmt, err := tx.Prepare("SELECT s.name, s.artist, COUNT(*) FROM song s JOIN listen l ON s.id = l.songID WHERE username = ? GROUP BY s.id ORDER BY COUNT(*) DESC;")
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
	stmt, err := tx.Prepare("SELECT a.name, a.artist, a.image, COUNT(*) FROM album a JOIN song s ON a.id = s.album JOIN listen l ON s.id = l.songID WHERE username = ? GROUP BY a.id ORDER BY COUNT(*) DESC;")
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

func GetDecadeBreakdown(tx *sql.Tx, username string) ([]model.DecadeBreakdown, error) {
	stmt, err := tx.Prepare("SELECT CONCAT(FLOOR(a.year / 10) * 10, '-', FLOOR(a.year / 10) * 10 + 9) AS decade, COUNT(*) FROM listen l JOIN song s ON l.songID = s.id JOIN album a ON s.album = a.id WHERE username = ? GROUP BY decade ORDER BY COUNT(*) DESC;")
	if err != nil {
		return nil, err
	}

	var results []model.DecadeBreakdown
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
		var decade string
		var count int
		err = rows.Scan(&decade, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.DecadeBreakdown{Decade: decade, Count: count})
	}

	return results, nil
}

func GetRawArtists(tx *sql.Tx, username string) ([]string, error) {
	stmt, err := tx.Prepare("SELECT artist FROM song s JOIN listen l ON s.id = l.songID WHERE username = ?;")
	if err != nil {
		return nil, err
	}

	var results []string
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
		var artist string
		err = rows.Scan(&artist)
		if err != nil {
			return nil, err
		}
		results = append(results, artist)
	}

	return results, nil
}
