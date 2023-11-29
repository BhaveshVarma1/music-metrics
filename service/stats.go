package service

import (
	"music-metrics/da"
	"music-metrics/model"
	"music-metrics/util"
	"sort"
	"strings"
	"time"
)

type StatsService interface {
	ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse
}

type AllStatsService struct{}

type AverageLengthService struct{}

type AveragePopularityService struct{}

type AverageYearService struct{}

type DecadeBreakdownService struct{}

type FirstTrackService struct{}

type HourBreakdownService struct{}

type MedianYearService struct{}

type ModeYearService struct{}

type PercentExplicitService struct{}

type TopAlbumsService struct{}

type TopAlbumsTimeService struct{}

type TopArtistsService struct{}

type TopArtistsTimeService struct{}

type TopTracksService struct{}

type TopTracksTimeService struct{}

type TotalMinutesService struct{}

type TotalTracksService struct{}

type UniqueAlbumsService struct{}

type UniqueArtistsService struct{}

type UniqueTracksService struct{}

type WeekDayBreakdownService struct{}

func (s AllStatsService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	var avgLengthService AverageLengthService
	var avgPopularityService AveragePopularityService
	var avgYearService AverageYearService
	var decadeBreakdownService DecadeBreakdownService
	var firstTrackService FirstTrackService
	var hourBreakdownService HourBreakdownService
	var medianYearService MedianYearService
	var modeYearService ModeYearService
	var percentExplicitService PercentExplicitService
	var topAlbumService TopAlbumsService
	var topAlbumTimeService TopAlbumsTimeService
	var topArtistService TopArtistsService
	var topArtistTimeService TopArtistsTimeService
	var topTrackService TopTracksService
	var topTrackTimeService TopTracksTimeService
	var totalMinutesService TotalMinutesService
	var totalTracksService TotalTracksService
	var uniqueAlbumsService UniqueAlbumsService
	var uniqueArtistsService UniqueArtistsService
	var uniqueTracksService UniqueTracksService
	var weekDayBreakdownService WeekDayBreakdownService

	totalTracks := totalTracksService.ExecuteService(username, startTime, endTime)
	if response, ok := totalTracks.(model.SingleIntResponse); ok {
		if response.Value < 1 {
			return "No tracks found for this time period."
		}
	}

	return model.AllStatsResponse{
		AverageLength:     avgLengthService.ExecuteService(username, startTime, endTime),
		AveragePopularity: avgPopularityService.ExecuteService(username, startTime, endTime),
		AverageYear:       avgYearService.ExecuteService(username, startTime, endTime),
		DecadeBreakdown:   decadeBreakdownService.ExecuteService(username, startTime, endTime),
		FirstTrack:        firstTrackService.ExecuteService(username, startTime, endTime),
		HourBreakdown:     hourBreakdownService.ExecuteService(username, startTime, endTime),
		MedianYear:        medianYearService.ExecuteService(username, startTime, endTime),
		ModeYear:          modeYearService.ExecuteService(username, startTime, endTime),
		PercentExplicit:   percentExplicitService.ExecuteService(username, startTime, endTime),
		TopAlbums:         topAlbumService.ExecuteService(username, startTime, endTime),
		TopAlbumsTime:     topAlbumTimeService.ExecuteService(username, startTime, endTime),
		TopArtists:        topArtistService.ExecuteService(username, startTime, endTime),
		TopArtistsTime:    topArtistTimeService.ExecuteService(username, startTime, endTime),
		TopTracks:         topTrackService.ExecuteService(username, startTime, endTime),
		TopTracksTime:     topTrackTimeService.ExecuteService(username, startTime, endTime),
		TotalMinutes:      totalMinutesService.ExecuteService(username, startTime, endTime),
		TotalTracks:       totalTracksService.ExecuteService(username, startTime, endTime),
		UniqueAlbums:      uniqueAlbumsService.ExecuteService(username, startTime, endTime),
		UniqueArtists:     uniqueArtistsService.ExecuteService(username, startTime, endTime),
		UniqueTracks:      uniqueTracksService.ExecuteService(username, startTime, endTime),
		WeekDayBreakdown:  weekDayBreakdownService.ExecuteService(username, startTime, endTime),
	}

}

func (s AverageLengthService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetAverageLength(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s AveragePopularityService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetAveragePopularityWithTracks(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if result == nil {
		return "No tracks found"
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.AveragePopularityResponse{Items: result}
}

func (s AverageYearService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetAverageYear(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s DecadeBreakdownService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetDecadeBreakdown(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.DecadeBreakdownResponse{Items: result}
}

func (s FirstTrackService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetFirstTrack(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return result
}

func (s HourBreakdownService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetRawTimestamps(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	hours := make([]int, 24)
	for _, timestamp := range result {
		timeObj := time.Unix(timestamp/1000, 0)
		hours[timeObj.Hour()]++
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.HourBreakdownResponse{Items: hours}
}

func (s MedianYearService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetRawYears(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	medianYear := result[len(result)/2]

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: medianYear}
}

func (s ModeYearService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetModeYears(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	// Calculate percentages
	total, err := da.GetTotalTracks(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	for i := range result {
		percent := float64(result[i].Count) / float64(total) * 100
		result[i].Count = int(percent)
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.ModeYearResponse{Items: result}
}

func (s PercentExplicitService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetPercentExplicit(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}

}

func (s TopAlbumsService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetTopAlbums(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopAlbumsResponse{Items: result}
}

func (s TopAlbumsTimeService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetTopAlbumsTime(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopAlbumsResponse{Items: result}
}

func (s TopArtistsService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetRawArtists(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
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

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	// Take only the first 1000 items
	if len(sortedArtists) > 1000 {
		sortedArtists = sortedArtists[:1000]
	}

	return model.TopArtistsResponse{Items: sortedArtists}
}

func (s TopArtistsTimeService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetRawArtists(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
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

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	// Take only the first 1000 items
	if len(toReturn) > 1000 {
		toReturn = toReturn[:1000]
	}

	return model.TopArtistsResponse{Items: toReturn}
}

func (s TopTracksService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetTopTracks(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopTracksResponse{Items: result}
}

func (s TopTracksTimeService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetTopTracksTime(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.TopTracksResponse{Items: result}
}

func (s TotalTracksService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetTotalTracks(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s TotalMinutesService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetTotalMinutes(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s UniqueAlbumsService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetUniqueAlbums(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s UniqueArtistsService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetRawArtists(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	count := 0
	var uniqueArtists []string
	for _, rawArtist := range result {
		artists := strings.Split(rawArtist.Artist, ";;")
		for _, artist := range artists {
			if !util.SliceContainsString(uniqueArtists, artist) {
				uniqueArtists = append(uniqueArtists, artist)
				count++
			}
		}
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: count}
}

func (s UniqueTracksService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetUniqueTracks(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.SingleIntResponse{Value: result}
}

func (s WeekDayBreakdownService) ExecuteService(username string, startTime int64, endTime int64) model.StatsResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return nil
	}

	result, err := da.GetRawTimestamps(tx, username, startTime, endTime)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return nil
		}
		return nil
	}

	days := make([]int, 7)
	for _, timestamp := range result {
		timeObj := time.Unix(timestamp/1000, 0)
		days[timeObj.Weekday()]++
	}

	if da.CommitAndClose(tx, db, true) != nil {
		return nil
	}

	return model.WeekDayBreakdownResponse{Items: days}
}
