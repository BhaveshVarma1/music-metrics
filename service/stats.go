package service

import (
	"music-metrics-back/dal"
)

func GetAverageYear(username string) int {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return -1
	}

	result, err := dal.GetAverageYear(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return -1
		}
		return -1
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return -1
	}

	return result

}
