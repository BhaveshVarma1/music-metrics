package da

import (
	"database/sql"
	"fmt"
	"math"
	"music-metrics/model"
)

func GetTopTracks(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.TopTrack, error) {
	stmt, err := tx.Prepare(SQL_TOP_TRACKS)
	if err != nil {
		return nil, err
	}

	var results []model.TopTrack
	rows, err := stmt.Query(username, startTime, endTime)
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
		var track string
		var trackId string
		var artist string
		var artistId string
		var image string
		var count int
		err = rows.Scan(&track, &trackId, &artist, &artistId, &image, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.TopTrack{
			Track:    track,
			Artist:   artist,
			Count:    count,
			TrackId:  trackId,
			ArtistId: artistId,
			Image:    image,
		})
	}

	return results, nil
}

func GetTopTracksTime(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.TopTrack, error) {
	stmt, err := tx.Prepare(SQL_TOP_TRACKS_TIME)
	if err != nil {
		return nil, err
	}

	var results []model.TopTrack
	rows, err := stmt.Query(username, startTime, endTime)
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
		var track string
		var trackId string
		var artist string
		var artistId string
		var image string
		var count int
		err = rows.Scan(&track, &trackId, &artist, &artistId, &image, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.TopTrack{
			Track:    track,
			Artist:   artist,
			Count:    count,
			TrackId:  trackId,
			ArtistId: artistId,
			Image:    image,
		})
	}

	return results, nil
}

func GetRawArtists(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.RawArtistTime, error) {
	stmt, err := tx.Prepare(SQL_RAW_ARTISTS)
	if err != nil {
		return nil, err
	}

	var results []model.RawArtistTime
	rows, err := stmt.Query(username, startTime, endTime)
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
		var artistId string
		var millis int
		err = rows.Scan(&artist, &artistId, &millis)
		if err != nil {
			return nil, err
		}
		results = append(results, model.RawArtistTime{
			Artist:   artist,
			Millis:   millis,
			ArtistId: artistId,
		})
	}

	return results, nil
}

func GetTopAlbums(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.TopAlbum, error) {
	stmt, err := tx.Prepare(SQL_TOP_ALBUMS)
	if err != nil {
		return nil, err
	}

	var results []model.TopAlbum
	rows, err := stmt.Query(username, startTime, endTime)
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
		var albumId string
		var artist string
		var artistId string
		var image string
		var count int
		err = rows.Scan(&album, &albumId, &artist, &artistId, &image, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.TopAlbum{
			Album:    album,
			Artist:   artist,
			Image:    image,
			Count:    count,
			AlbumId:  albumId,
			ArtistId: artistId,
		})
	}

	return results, nil
}

func GetTopAlbumsTime(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.TopAlbum, error) {
	stmt, err := tx.Prepare(SQL_TOP_ALBUMS_TIME)
	if err != nil {
		return nil, err
	}

	var results []model.TopAlbum
	rows, err := stmt.Query(username, startTime, endTime)
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
		var albumId string
		var artist string
		var artistId string
		var image string
		var count int
		err = rows.Scan(&album, &albumId, &artist, &artistId, &image, &count)
		if err != nil {
			return nil, err
		}
		results = append(results, model.TopAlbum{
			Album:    album,
			Artist:   artist,
			Image:    image,
			Count:    count,
			AlbumId:  albumId,
			ArtistId: artistId,
		})
	}

	return results, nil
}

func GetAverageYear(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_AVG_YEAR)
	if err != nil {
		return 0, err
	}

	var result float32
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return int(math.Round(float64(result))), nil
}

func GetDecadeBreakdown(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.DecadeBreakdown, error) {
	stmt, err := tx.Prepare(SQL_DECADE_BREAKDOWN)
	if err != nil {
		return nil, err
	}

	var results []model.DecadeBreakdown
	rows, err := stmt.Query(username, startTime, endTime)
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

func GetAverageLength(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_AVG_LENGTH)
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetPercentExplicit(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_PERCENT_EXPLICIT)
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetTotalTracks(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_TOTAL_TRACKS)
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetUniqueTracks(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_UNIQUE_TRACKS)
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetUniqueAlbums(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_UNIQUE_ALBUMS)
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetModeYears(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.ModeYear, error) {
	stmt, err := tx.Prepare(SQL_MODE_YEARS)
	if err != nil {
		return nil, err
	}

	var results []model.ModeYear
	rows, err := stmt.Query(username, startTime, endTime)
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

func GetRawYears(tx *sql.Tx, username string, startTime int64, endTime int64) ([]int, error) {
	stmt, err := tx.Prepare(SQL_RAW_YEARS)
	if err != nil {
		return nil, err
	}

	var results []int
	rows, err := stmt.Query(username, startTime, endTime)
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

func GetRawTimestamps(tx *sql.Tx, username string, startTime int64, endTime int64) ([]int64, error) {
	stmt, err := tx.Prepare(SQL_RAW_TIMESTAMPS)
	if err != nil {
		return nil, err
	}

	var results []int64
	rows, err := stmt.Query(username, startTime, endTime)
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

func GetAveragePopularityWithTracks(tx *sql.Tx, username string, startTime int64, endTime int64) ([]model.PopularityObject, error) {
	stmt, err := tx.Prepare(SQL_AVG_POPULARITY)
	if err != nil {
		return nil, err
	}

	var results []model.PopularityObject
	rows, err := stmt.Query(username, username, startTime, endTime)
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
		var track string
		var trackId string
		var artist string
		var artistId string
		var popularity int
		err = rows.Scan(&track, &trackId, &artist, &artistId, &popularity)
		if err != nil {
			return nil, err
		}
		results = append(results, model.PopularityObject{
			Track:      track,
			Artist:     artist,
			Popularity: popularity,
			TrackId:    trackId,
			ArtistId:   artistId,
		})
	}

	return results, nil
}

func GetTotalMinutes(tx *sql.Tx, username string, startTime int64, endTime int64) (int, error) {
	stmt, err := tx.Prepare(SQL_TOTAL_MINUTES)
	if err != nil {
		return 0, err
	}

	var result int
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result)
	if err != nil {
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetFirstTrack(tx *sql.Tx, username string, startTime int64, endTime int64) (model.FirstTrackResponse, error) {
	stmt, err := tx.Prepare(SQL_FIRST_TRACK)
	if err != nil {
		return model.FirstTrackResponse{}, err
	}

	var result model.FirstTrackResponse
	err = stmt.QueryRow(username, startTime, endTime).Scan(&result.Track, &result.TrackId, &result.Artist, &result.ArtistId, &result.Image, &result.Timestamp)
	if err != nil {
		return model.FirstTrackResponse{}, err
	}

	err = stmt.Close()
	if err != nil {
		return model.FirstTrackResponse{}, err
	}
	return result, nil
}
