package da

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Db() *sql.DB {
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		return nil
	}
	database, err := sql.Open("mysql", "pratt:"+password+"@tcp(musicmetrics.app:3306)/mm")
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return nil
	}
	err = database.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		return nil
	}
	return database

}

func DbClose(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func BeginTX() (*sql.Tx, *sql.DB, error) {
	db := Db()
	if db == nil {
		return nil, nil, errors.New("db is nil")
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, err
	}
	return tx, db, nil
}

func CommitAndClose(tx *sql.Tx, db *sql.DB, commit bool) error {
	if commit {
		err := tx.Commit()
		if err != nil {
			return err
		}
	} else {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	}
	err := DbClose(db)
	if err != nil {
		return err
	}
	return nil
}
