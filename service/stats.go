package service

import (
	"music-metrics/dal"
	"music-metrics/model"
)

type StatsService interface {
	ExecuteService(username string) model.StatsResponse
}

type GetAverageYearService struct{}

type GetSongCountsService struct{}

type GetTopAlbumsService struct{}

type GetDecadeBreakdownService struct{}

func (s GetAverageYearService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetAverageYear(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.AverageYearResponse{Success: true, AverageYear: result}
}

func (s GetSongCountsService) ExecuteService(username string) model.StatsResponse {

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

	return model.SongCountsResponse{Success: true, SongCounts: result}
}

func (s GetTopAlbumsService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetTopAlbums(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopAlbumsResponse{Success: true, TopAlbums: result}
}

func (s GetDecadeBreakdownService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetDecadeBreakdown(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.DecadeBreakdownResponse{Success: true, DecadeBreakdowns: result}
}
