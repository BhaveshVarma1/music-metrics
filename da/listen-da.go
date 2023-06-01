package da

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateListen(tx *sql.Tx, listen model.ListenBean) error {
	_, err := tx.Exec("INSERT INTO listen (username, timestamp, songID) VALUES(?,?,?);", listen.Username, listen.Timestamp, listen.SongId)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveListen(tx *sql.Tx, username string, timestamp int64) (model.ListenBean, error) {
	rows, err := tx.Query("SELECT * FROM listen WHERE username = ? AND timestamp = ?;", username, timestamp)
	if err != nil {
		return model.ListenBean{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var listen model.ListenBean
	for rows.Next() {
		err = rows.Scan(&listen.Username, &listen.Timestamp, &listen.SongId)
		if err != nil {
			return model.ListenBean{}, err
		}
		return listen, nil
	}
	return model.ListenBean{}, nil
}

func RetrieveAllListensForUser(tx *sql.Tx, username string) ([]model.ListenBean, error) {
	rows, err := tx.Query("SELECT * FROM listen WHERE username = ?;", username)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var listens []model.ListenBean
	for rows.Next() {
		var listen model.ListenBean
		err = rows.Scan(&listen.Username, &listen.Timestamp, &listen.SongId)
		if err != nil {
			return nil, err
		}
		listens = append(listens, listen)
	}
	return listens, nil
}

func RetrieveAllTimestampsForUser(tx *sql.Tx, username string) ([]int64, error) {
	rows, err := tx.Query("SELECT timestamp FROM listen WHERE username = ?;", username)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var timestamps []int64
	for rows.Next() {
		var timestamp int64
		err = rows.Scan(&timestamp)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, timestamp)
	}
	return timestamps, nil
}

func DeleteListen(tx *sql.Tx, username string, timestamp int64) error {
	_, err := tx.Exec("DELETE FROM listen WHERE username = ? AND timestamp = ?;", username, timestamp)
	if err != nil {
		return err
	}
	return nil
}

func DeleteListenByUsername(tx *sql.Tx, username string) error {
	_, err := tx.Exec("DELETE FROM listen WHERE username = ?;", username)
	if err != nil {
		return err
	}
	return nil
}

func GetMostRecentListen(tx *sql.Tx, username string) (model.ListenBean, error) {
	rows, err := tx.Query("SELECT * FROM listen WHERE username = ? ORDER BY timestamp DESC LIMIT 1;", username)
	if err != nil {
		return model.ListenBean{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var listen model.ListenBean
	for rows.Next() {
		err = rows.Scan(&listen.Username, &listen.Timestamp, &listen.SongId)
		if err != nil {
			return model.ListenBean{}, err
		}
		return listen, nil
	}
	return model.ListenBean{}, nil
}

func ClearListen(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM listen;")
	if err != nil {
		return err
	}
	return nil
}

func HasTimestamp(tx *sql.Tx, username string, timestamp int64) (bool, error) {
	rows, err := tx.Query("SELECT * FROM listen WHERE username = ? AND timestamp = ?;", username, timestamp)
	if err != nil {
		return false, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
