package dal

import (
	"database/sql"
	"fmt"
	"music-metrics-back/model"
)

func CreateUser(tx *sql.Tx, user model.User) error {
	_, err := tx.Exec("INSERT INTO user (username, displayName, email, code, refresh, access) VALUES (?, ?, ?, ?, ?, ?);",
		user.Username, user.DisplayName, user.Email, user.Code, user.Refresh, user.Access)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveUser(tx *sql.Tx, username string) (model.User, error) {
	rows, err := tx.Query("SELECT * FROM user WHERE username = ?;", username)
	if err != nil {
		return model.User{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.DisplayName, &user.Email, &user.Code, &user.Refresh, &user.Access)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	}
	return model.User{}, nil
}

func UpdateUser(tx *sql.Tx, user model.User) error {
	_, err := tx.Exec("UPDATE user SET displayName = ?, email = ?, code = ?, refresh = ?, access = ? WHERE username = ?;",
		user.DisplayName, user.Email, user.Code, user.Refresh, user.Access, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(tx *sql.Tx, username string) error {
	_, err := tx.Exec("DELETE FROM user WHERE username = ?;", username)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveUserByCode(tx *sql.Tx, code string) (model.User, error) {
	rows, err := tx.Query("SELECT * FROM user WHERE code = ?;", code)
	if err != nil {
		return model.User{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.DisplayName, &user.Email, &user.Code, &user.Refresh, &user.Access)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	}
	return model.User{}, nil
}

func ClearUsers(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM user;")
	if err != nil {
		return err
	}
	return nil
}
