package service

import (
	"music-metrics-back/dal"
	"music-metrics-back/model"
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

func GetSongCounts(username string) []model.SongCount {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetSongCounts(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return result

}
