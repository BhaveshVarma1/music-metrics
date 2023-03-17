package dal

import (
	"database/sql"
	"fmt"
	"music-metrics-back/model"
)

func CreateListen(tx *sql.Tx, listen model.Listen) error {
	_, err := tx.Exec("INSERT INTO listen (username, timestamp, songID) VALUES(?,?,?);", listen.Username, listen.Timestamp, listen.SongId)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveListen(tx *sql.Tx, username string, timestamp string) (model.Listen, error) {
	rows, err := tx.Query("SELECT * FROM listen WHERE username = ? AND timestamp = ?;", username, timestamp)
	if err != nil {
		return model.Listen{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var listen model.Listen
	for rows.Next() {
		err = rows.Scan(&listen.Username, &listen.Timestamp, &listen.SongId)
		if err != nil {
			return model.Listen{}, err
		}
		return listen, nil
	}
	return model.Listen{}, nil
}

func DeleteListen(tx *sql.Tx, username string, timestamp string) error {
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

func ClearListen(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM listen;")
	if err != nil {
		return err
	}
	return nil
}
