package service

import (
	"fmt"
	"music-metrics/da"
)

func Delete(username string) {

	// Instantiate database connection
	tx, db, err := da.BeginTX()
	if err != nil {
		fmt.Println("Error beginning transaction in delete service: ", err)
		return
	}

	// Delete user
	err = da.DeleteUser(tx, username)
	if err != nil {
		fmt.Println("Error deleting user in delete service: ", err)
		return
	}

	// Commit changes and close the connection
	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing and closing in delete service: ", err)
		return
	}
}
