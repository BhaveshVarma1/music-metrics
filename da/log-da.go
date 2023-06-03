package da

import (
	"database/sql"
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
