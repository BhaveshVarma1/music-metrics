package service

import (
	"music-metrics/dal"
	"music-metrics/model"
	"sort"
	"strings"
	"time"
)

type StatsService interface {
	ExecuteService(username string) model.StatsResponse
}

type AverageLengthService struct{}

type AveragePopularityService struct{}

type AverageYearService struct{}

type DecadeBreakdownService struct{}

type HourBreakdownService struct{}

type MedianYearService struct{}

type ModeYearService struct{}

type PercentExplicitService struct{}

type TopAlbumsService struct{}

type TopAlbumsTimeService struct{}

type TopArtistsService struct{}

type TopArtistsTimeService struct{}

type TopSongsService struct{}

type TopSongsTimeService struct{}

type TotalSongsService struct{}

type UniqueAlbumsService struct{}

type UniqueArtistsService struct{}

type UniqueSongsService struct{}

type WeekDayBreakdownService struct{}

func (s AverageLengthService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetAverageLength(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s AveragePopularityService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetAveragePopularityWithSongs(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if result == nil {
		return "No songs found"
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.AveragePopularityResponse{Items: result}
}

func (s AverageYearService) ExecuteService(username string) model.StatsResponse {

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

	return model.SingleIntResponse{Value: result}
}

func (s DecadeBreakdownService) ExecuteService(username string) model.StatsResponse {

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

	return model.DecadeBreakdownResponse{Items: result}
}

func (s HourBreakdownService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetRawTimestamps(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	hours := make([]int, 24)
	for _, timestamp := range result {
		timeObj := time.Unix(timestamp/1000, 0)
		hours[timeObj.Hour()]++
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.HourBreakdownResponse{Items: hours}
}

func (s MedianYearService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetRawYears(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	medianYear := result[len(result)/2]

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: medianYear}
}

func (s ModeYearService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetModeYears(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	// Calculate percentages
	total, err := dal.GetTotalSongs(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	for i := range result {
		percent := float64(result[i].Count) / float64(total) * 100
		result[i].Count = int(percent)
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.ModeYearResponse{Items: result}
}

func (s PercentExplicitService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetPercentExplicit(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}

}

func (s TopAlbumsService) ExecuteService(username string) model.StatsResponse {

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

	return model.TopAlbumsResponse{Items: result}
}

func (s TopAlbumsTimeService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetTopAlbumsTime(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopAlbumsResponse{Items: result}
}

func (s TopArtistsService) ExecuteService(username string) model.StatsResponse {

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

	// Map of artist name to artist id
	artistsWithIds := make(map[string]string)
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist.Artist, ";;")
		ids := strings.Split(rawArtist.ArtistId, ";;")
		for i, artist := range artists {
			artistsWithIds[artist] = ids[i]
		}
	}

	// Create a proper map of top artists and counts since they are stored with ';;' in the db
	topArtists := make(map[string]int)
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist.Artist, ";;")
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
		sortedArtists = append(sortedArtists, model.TopArtist{Artist: artist, Count: count, ArtistId: artistsWithIds[artist]})
	}
	sort.Slice(sortedArtists, func(i, j int) bool {
		return sortedArtists[i].Count > sortedArtists[j].Count
	})

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopArtistsResponse{Items: sortedArtists}
}

func (s TopArtistsTimeService) ExecuteService(username string) model.StatsResponse {

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

	// Map of artist name to artist id
	artistsWithIds := make(map[string]string)
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist.Artist, ";;")
		ids := strings.Split(rawArtist.ArtistId, ";;")
		for i, artist := range artists {
			artistsWithIds[artist] = ids[i]
		}
	}

	counts := make(map[string]int)
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist.Artist, ";;")
		for _, artist := range artists {
			if _, ok := counts[artist]; ok {
				counts[artist] += rawArtist.Millis
			} else {
				counts[artist] = rawArtist.Millis
			}
		}
	}

	var toReturn []model.TopArtist
	for k, v := range counts {
		toReturn = append(toReturn, model.TopArtist{Artist: k, Count: v / 1000, ArtistId: artistsWithIds[k]})
	}

	sort.Slice(toReturn, func(i, j int) bool {
		return toReturn[i].Count > toReturn[j].Count
	})

	return model.TopArtistsResponse{Items: toReturn}
}

func (s TopSongsService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetTopSongs(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopSongsResponse{Items: result}
}

func (s TopSongsTimeService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetTopSongsTime(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopSongsResponse{Items: result}
}

func (s TotalSongsService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetTotalSongs(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s UniqueAlbumsService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetUniqueAlbums(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s UniqueArtistsService) ExecuteService(username string) model.StatsResponse {

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

	count := 0
	var uniqueArtists []string
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist.Artist, ";;")
		for _, artist := range artists {
			if !SliceContainsString(uniqueArtists, artist) {
				uniqueArtists = append(uniqueArtists, artist)
				count++
			}
		}
	}

	return model.SingleIntResponse{Value: count}
}

func (s UniqueSongsService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetUniqueSongs(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s WeekDayBreakdownService) ExecuteService(username string) model.StatsResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return nil
	}

	result, err := dal.GetRawTimestamps(tx, username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	days := make([]int, 7)
	for _, timestamp := range result {
		timeObj := time.Unix(timestamp/1000, 0)
		days[timeObj.Weekday()]++
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.WeekDayBreakdownResponse{Items: days}
}
