package dal

import (
	"database/sql"
	"fmt"
	"math"
	"music-metrics/model"
)

func GetTopSongs(tx *sql.Tx, username string) ([]model.TopSong, error) {
	stmt, err := tx.Prepare("SELECT s.name, s.artist, COUNT(*) FROM song s JOIN listen l ON s.id = l.songID WHERE username = ? GROUP BY s.id ORDER BY COUNT(*) DESC;")
	if err != nil {
		return nil, err
	}

	var results []model.TopSong
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
		results = append(results, model.TopSong{Song: song, Artist: artist, Count: count})
	}

	return results, nil
}

func GetTopSongsTime(tx *sql.Tx, username string) ([]model.TopSong, error) {
	stmt, err := tx.Prepare("SELECT s.name, s.artist, ROUND(COUNT(*) * s.duration / 1000) AS time FROM listen l JOIN song s ON l.songID = s.id WHERE username = ? GROUP BY s.id ORDER BY time DESC LIMIT 1000;")
	if err != nil {
		return nil, err
	}

	var results []model.TopSong
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
		results = append(results, model.TopSong{Song: song, Artist: artist, Count: count})
	}

	return results, nil
}

func GetRawArtists(tx *sql.Tx, username string) ([]model.RawArtistTime, error) {
	stmt, err := tx.Prepare("SELECT artist, duration FROM song s JOIN listen l ON s.id = l.songID WHERE username = ?;")
	if err != nil {
		return nil, err
	}

	var results []model.RawArtistTime
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
		var millis int
		err = rows.Scan(&artist, &millis)
		if err != nil {
			return nil, err
		}
		results = append(results, model.RawArtistTime{Artist: artist, Millis: millis})
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

func GetTopAlbumsTime(tx *sql.Tx, username string) ([]model.TopAlbum, error) {
	stmt, err := tx.Prepare("SELECT a.name, a.artist, a.image, ROUND(SUM(x.time) / 1000) FROM (SELECT s.album, (COUNT(*) * s.duration) AS time FROM listen l JOIN song s ON l.songID = s.id WHERE username = ? GROUP BY s.id ORDER BY time DESC) AS x JOIN album a ON x.album = a.id GROUP BY x.album ORDER BY SUM(x.time) DESC LIMIT 1000;")
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

func GetDecadeBreakdown(tx *sql.Tx, username string) ([]model.DecadeBreakdown, error) {
	stmt, err := tx.Prepare("SELECT CONCAT(FLOOR(a.year / 10) * 10, 's') AS decade, COUNT(*) FROM listen l JOIN song s ON l.songID = s.id JOIN album a ON s.album = a.id WHERE username = ? GROUP BY decade ORDER BY COUNT(*) DESC;")
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

func GetAverageLength(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT ROUND(AVG(s.duration) / 1000, 0) FROM listen l JOIN song s ON l.songID = s.id WHERE username = ?;")
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetPercentExplicit(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT ROUND(100 * AVG(s.explicit = 1), 0) FROM listen l JOIN song s ON l.songID = s.id WHERE l.username = ?;")
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetTotalSongs(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT COUNT(*) FROM listen l WHERE username = ?;")
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetUniqueSongs(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT COUNT(DISTINCT s.id) FROM listen l JOIN song s ON l.songID = s.id WHERE username = ?;")
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetUniqueAlbums(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT COUNT(DISTINCT a.id) FROM listen l JOIN song s ON l.songID = s.id JOIN album a ON s.album = a.id WHERE username = ?;")
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetModeYears(tx *sql.Tx, username string) ([]model.ModeYear, error) {
	stmt, err := tx.Prepare("SELECT a.year, COUNT(*) FROM listen l JOIN song s ON l.songID = s.id JOIN album a ON s.album = a.id WHERE username = ? GROUP BY year ORDER BY COUNT(*) DESC LIMIT 3;")
	if err != nil {
		return nil, err
	}

	var results []model.ModeYear
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
		var year int
		var count int
		err = rows.Scan(&year, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.ModeYear{Year: year, Count: count})
	}

	return results, nil
}

func GetRawYears(tx *sql.Tx, username string) ([]int, error) {
	stmt, err := tx.Prepare("SELECT a.year FROM listen l JOIN song s ON l.songID = s.id JOIN album a ON s.album = a.id WHERE username = ? ORDER BY a.year;")
	if err != nil {
		return nil, err
	}

	var results []int
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
		var year int
		err = rows.Scan(&year)
		if err != nil {
			return nil, err
		}
		results = append(results, year)
	}

	return results, nil
}

func GetRawTimestamps(tx *sql.Tx, username string) ([]int64, error) {
	stmt, err := tx.Prepare("SELECT l.timestamp FROM listen l WHERE username = ? ORDER BY l.timestamp;")
	if err != nil {
		return nil, err
	}

	var results []int64
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
		var timestamp int64
		err = rows.Scan(&timestamp)
		if err != nil {
			return nil, err
		}
		results = append(results, timestamp)
	}

	return results, nil
}

func GetAveragePopularityWithSongs(tx *sql.Tx, username string) ([]model.PopularityObject, error) {
	stmt, err := tx.Prepare("SELECT s.name, s.artist, s.popularity FROM listen l JOIN song s ON l.songID = s.id WHERE l.username = ? AND s.popularity = ROUND((SELECT AVG(s.popularity) FROM listen l JOIN song s ON l.songID = s.id WHERE l.username = ?), 0) GROUP BY s.id LIMIT 3;")
	if err != nil {
		return nil, err
	}

	var results []model.PopularityObject
	rows, err := stmt.Query(username, username)
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
		var popularity int
		err = rows.Scan(&song, &artist, &popularity)
		if err != nil {
			return nil, err
		}
		results = append(results, model.PopularityObject{Song: song, Artist: artist, Popularity: popularity})
	}

	return results, nil
}
