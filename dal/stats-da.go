package dal

import "database/sql"

func GetAverageYear(tx *sql.Tx, username string) (int, error) {
	stmt, err := tx.Prepare("SELECT AVG(year) FROM song INNER JOIN listen ON song.id = listen.songID WHERE username = ?;")
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
