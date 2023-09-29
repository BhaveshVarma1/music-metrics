package service

import (
	"fmt"
	"testing"
	"time"
)

func GetUsername() string {
	return "prattnj"
}

func GetStartTime() int64 {

	// April 1st, 2023 (Utah): 1680328800000
	// May 1st, 2023 (Utah): 1682920800000

	return 0
}

func GetEndTime() int64 {
	return time.Now().UnixMilli()
}

func TestAllStats(t *testing.T) {

	var service AllStatsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestAverageLength(t *testing.T) {

	var service AverageLengthService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestAveragePopularity(t *testing.T) {

	var service AveragePopularityService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestAverageYear(t *testing.T) {

	var service AverageYearService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestDecadeBreakdown(t *testing.T) {

	var service DecadeBreakdownService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestFirstTrack(t *testing.T) {

	var service FirstTrackService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestHourBreakdown(t *testing.T) {

	var service HourBreakdownService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestMedianYear(t *testing.T) {

	var service MedianYearService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestModeYear(t *testing.T) {

	var service ModeYearService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestPercentExplicit(t *testing.T) {

	var service PercentExplicitService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopAlbums(t *testing.T) {

	var service TopAlbumsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopAlbumsTime(t *testing.T) {

	var service TopAlbumsTimeService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopArtists(t *testing.T) {

	var service TopArtistsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopArtistsTime(t *testing.T) {

	var service TopArtistsTimeService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopTracks(t *testing.T) {

	var service TopTracksService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopTracksTime(t *testing.T) {

	var service TopTracksTimeService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTotalMinutes(t *testing.T) {

	var service TotalMinutesService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTotalTracks(t *testing.T) {

	var service TotalTracksService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestUniqueAlbums(t *testing.T) {

	var service UniqueAlbumsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestGetUniqueArtists(t *testing.T) {

	var service UniqueArtistsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)

}

func TestUniqueTracks(t *testing.T) {

	var service UniqueTracksService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestWeekDayBreakdown(t *testing.T) {

	var service WeekDayBreakdownService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}
