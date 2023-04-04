package service

import (
	"music-metrics/dal"
	"music-metrics/model"
	"sort"
	"strings"
)

type StatsService interface {
	ExecuteService(username string) model.StatsResponse
}

type GetAverageYearService struct{}

type GetSongCountsService struct{}

type GetTopAlbumsService struct{}

type GetDecadeBreakdownService struct{}

type GetTopArtistsService struct{}

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

func (s GetTopArtistsService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetRawArtists(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	// Create a proper map of top artists and counts since they are stored with ';;' in the db
	topArtists := make(map[string]int)
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist, ";;")
		for _, artist := range artists {
			if _, ok := topArtists[artist]; ok {
				topArtists[artist]++
			} else {
				topArtists[artist] = 1
			}
		}
	}

	// Sort the map by descending count
	sortedArtists := make([]model.TopArtist, 0)
	for artist, count := range topArtists {
		sortedArtists = append(sortedArtists, model.TopArtist{Artist: artist, Count: count})
	}
	sort.Slice(sortedArtists, func(i, j int) bool {
		return sortedArtists[i].Count > sortedArtists[j].Count
	})

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopArtistsResponse{Success: true, TopArtists: sortedArtists}
}
