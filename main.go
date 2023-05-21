package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics/handler"
	"music-metrics/service"
)

func main() {

	buildPath := "public/build"

	e := echo.New()

	// Handle WebSocket connections
	e.GET("/ws", handler.HandleWebsocket)

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	var allStatsService service.AllStatsService
	/*var avgLengthService service.AverageLengthService
	var avgPopularityService service.AveragePopularityService
	var avgYearService service.AverageYearService
	var decadeBreakdownService service.DecadeBreakdownService
	var hourBreakdownService service.HourBreakdownService
	var medianYearService service.MedianYearService
	var modeYearService service.ModeYearService
	var percentExplicitService service.PercentExplicitService
	var topAlbumService service.TopAlbumsService
	var topAlbumTimeService service.TopAlbumsTimeService
	var topArtistService service.TopArtistsService
	var topArtistTimeService service.TopArtistsTimeService
	var topSongService service.TopSongsService
	var topSongTimeService service.TopSongsTimeService
	var totalMinutesService service.TotalMinutesService
	var totalSongsService service.TotalSongsService
	var uniqueAlbumsService service.UniqueAlbumsService
	var uniqueArtistsService service.UniqueArtistsService
	var uniqueSongsService service.UniqueSongsService
	var weekDayBreakdownService service.WeekDayBreakdownService*/

	// API ENDPOINTS
	e.POST("/api/v1/updateCode", handler.HandleUpdateCode)
	e.GET("/api/v1/allStats/:username/:range", handler.StatsHandler(allStatsService))
	/*e.GET("/api/v1/averageLength/:username", handler.StatsHandler(avgLengthService))
	e.GET("/api/v1/averagePopularity/:username", handler.StatsHandler(avgPopularityService))
	e.GET("/api/v1/averageYear/:username", handler.StatsHandler(avgYearService))
	e.GET("/api/v1/decadeBreakdown/:username", handler.StatsHandler(decadeBreakdownService))
	e.GET("/api/v1/hourBreakdown/:username", handler.StatsHandler(hourBreakdownService))
	e.GET("/api/v1/medianYear/:username", handler.StatsHandler(medianYearService))
	e.GET("/api/v1/modeYear/:username", handler.StatsHandler(modeYearService))
	e.GET("/api/v1/percentExplicit/:username", handler.StatsHandler(percentExplicitService))
	e.GET("/api/v1/topAlbums/:username", handler.StatsHandler(topAlbumService))
	e.GET("/api/v1/topAlbumsTime/:username", handler.StatsHandler(topAlbumTimeService))
	e.GET("/api/v1/topArtists/:username", handler.StatsHandler(topArtistService))
	e.GET("/api/v1/topArtistsTime/:username", handler.StatsHandler(topArtistTimeService))
	e.GET("/api/v1/topSongs/:username", handler.StatsHandler(topSongService))
	e.GET("/api/v1/topSongsTime/:username", handler.StatsHandler(topSongTimeService))
	e.GET("/api/v1/totalSongs/:username", handler.StatsHandler(totalSongsService))
	e.GET("/api/v1/totalMinutes/:username", handler.StatsHandler(totalMinutesService))
	e.GET("/api/v1/uniqueAlbums/:username", handler.StatsHandler(uniqueAlbumsService))
	e.GET("/api/v1/uniqueArtists/:username", handler.StatsHandler(uniqueArtistsService))
	e.GET("/api/v1/uniqueSongs/:username", handler.StatsHandler(uniqueSongsService))
	e.GET("/api/v1/weekDayBreakdown/:username", handler.StatsHandler(weekDayBreakdownService))*/

	// STATIC / REACT FILES
	e.GET("/static/*", func(c echo.Context) error {
		return c.File(buildPath + c.Request().URL.Path)
	})

	e.GET("/manifest.json", func(c echo.Context) error {
		return c.File(buildPath + "/manifest.json")
	})

	e.GET("/favicon.ico", func(c echo.Context) error {
		return c.File(buildPath + "/favicon.ico")
	})

	e.GET("/*", func(c echo.Context) error {
		return c.File(buildPath + "/index.html")
	})

	e.Logger.Fatal(e.Start(":3000"))

}
