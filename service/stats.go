package service

import (
	"fmt"
	"music-metrics-back/dal"
)

func GetAverageYear(username string) int {

	tx, db, err := dal.BeginTX()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	result, err := dal.GetAverageYear(tx, username)
	if err != nil {
		fmt.Println(err)
		if dal.CommitAndClose(tx, db, false) != nil {
			fmt.Println("yo")
			return -1
		}
		return -1
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		fmt.Println("yo yo")
		return -1
	}

	return result

}
