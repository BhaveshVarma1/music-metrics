package da

import (
	"database/sql"
	"fmt"
	"music-metrics/model"
)

func CreateUser(tx *sql.Tx, user model.UserBean) error {
	_, err := tx.Exec("INSERT INTO user (username, displayName, email, refresh, timestamp) VALUES (?, ?, ?, ?, ?);",
		user.Username, user.DisplayName, user.Email, user.Refresh, user.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveUser(tx *sql.Tx, username string) (model.UserBean, error) {
	rows, err := tx.Query("SELECT * FROM user WHERE username = ?;", username)
	if err != nil {
		return model.UserBean{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var user model.UserBean
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.DisplayName, &user.Email, &user.Refresh, &user.Timestamp)
		if err != nil {
			return model.UserBean{}, err
		}
		return user, nil
	}
	return model.UserBean{}, nil
}

func UpdateUser(tx *sql.Tx, user model.UserBean) error {
	_, err := tx.Exec("UPDATE user SET displayName = ?, email = ?, refresh = ?, timestamp = ? WHERE username = ?;",
		user.DisplayName, user.Email, user.Refresh, user.Timestamp, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAllUsers(tx *sql.Tx) ([]model.UserBean, error) {
	rows, err := tx.Query("SELECT * FROM user;")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var users []model.UserBean
	for rows.Next() {
		var user model.UserBean
		err = rows.Scan(&user.Username, &user.DisplayName, &user.Email, &user.Refresh, &user.Timestamp)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func DeleteUser(tx *sql.Tx, username string) error {
	_, err := tx.Exec("DELETE FROM user WHERE username = ?;", username)
	if err != nil {
		return err
	}
	return nil
}

func ClearUsers(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM user;")
	if err != nil {
		return err
	}
	return nil
}
