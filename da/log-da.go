package da

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateLog(tx *sql.Tx, log *model.LogBean) error {
	_, err := tx.Exec("INSERT INTO log (`username`, `timestamp`, `action`, `ip`) VALUES (?, ?, ?, ?);",
		log.Username, log.Timestamp, log.Action, log.IP)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveLastAction(tx *sql.Tx) (string, int64, error) {
	rows, err := tx.Query("SELECT `username`, `timestamp` FROM log WHERE `username` != 'prattnj' ORDER BY timestamp DESC LIMIT 1;")
	if err != nil {
		return "", 0, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var log model.LogBean
	for rows.Next() {
		err = rows.Scan(&log.Username, &log.Timestamp)
		if err != nil {
			return "", 0, err
		}
		return log.Username, log.Timestamp, nil
	}
	return "", 0, nil
}
