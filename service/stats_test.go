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

	return 1682920800000
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

func TestTopSongs(t *testing.T) {

	var service TopSongsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTopSongsTime(t *testing.T) {

	var service TopSongsTimeService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestTotalSongs(t *testing.T) {

	var service TotalSongsService
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

func TestUniqueSongs(t *testing.T) {

	var service UniqueSongsService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}

func TestWeekDayBreakdown(t *testing.T) {

	var service WeekDayBreakdownService
	result := service.ExecuteService(GetUsername(), GetStartTime(), GetEndTime())
	fmt.Println(result)
}
