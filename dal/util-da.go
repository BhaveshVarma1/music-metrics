package dal

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Db() *sql.DB {
	password := GetPassword()
	if password == "" {
		return nil
	}
	database, err := sql.Open("mysql", "pratt:"+password+"@tcp(localhost:3306)/mm")
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

func GetPassword() string {
	file, err := os.Open("nogit.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
