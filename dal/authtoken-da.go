package dal

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateAuthToken(tx *sql.Tx, token model.AuthToken) error {
	_, err := tx.Exec("INSERT INTO authtoken (token, username, expiration) VALUES(?,?,?);", token.Token, token.Username, token.Expiration)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAuthToken(tx *sql.Tx, token string) (model.AuthToken, error) {
	rows, err := tx.Query("SELECT * FROM authtoken WHERE token = ?;", token)
	if err != nil {
		return model.AuthToken{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var authToken model.AuthToken
	for rows.Next() {
		err = rows.Scan(&authToken.Token, &authToken.Username, &authToken.Expiration)
		if err != nil {
			return model.AuthToken{}, err
		}
		return authToken, nil
	}
	return model.AuthToken{}, nil
}

func RetrieveAuthTokenByUsername(tx *sql.Tx, username string) (model.AuthToken, error) {
	rows, err := tx.Query("SELECT * FROM authtoken WHERE username = ?;", username)
	if err != nil {
		return model.AuthToken{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var authToken model.AuthToken
	for rows.Next() {
		err = rows.Scan(&authToken.Token, &authToken.Username, &authToken.Expiration)
		if err != nil {
			return model.AuthToken{}, err
		}
		return authToken, nil
	}
	return model.AuthToken{}, nil
}

func DeleteAuthToken(tx *sql.Tx, token string) error {
	_, err := tx.Exec("DELETE FROM authtoken WHERE token = ?;", token)
	if err != nil {
		return err
	}
	return nil
}

func ClearAuthTokens(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM authtoken;")
	if err != nil {
		return err
	}
	return nil
}
