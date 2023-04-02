package main

import (
	"fmt"
	"music-metrics/dal"
)

func main() {

	db := dal.Db()
	if db == nil {
		return
	}
	tx, err := db.Begin()
	if err != nil {
		return
	}

	//err = dal.ClearUsers(tx)
	//err = dal.ClearAuthTokens(tx)
	//err = dal.ClearSongs(tx)
	//err = dal.ClearListen(tx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return
	}

	//fmt.Println("Database cleared.")

}
