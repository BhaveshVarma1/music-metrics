package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics/handler"
	"music-metrics/service"
	"net/http"
)

func main() {

	buildPath := "public/build"

	e := echo.New()

	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Handle WebSocket connections
	e.GET("/ws", func(c echo.Context) error {
		fmt.Println("Websocket connection entered.")
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		// Handle WebSocket messages
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				return err
			}

			// Handle the received message
			// ...
			fmt.Println(string(message))

			// Send response
			err = conn.WriteMessage(messageType, message)
			if err != nil {
				return err
			}
		}
	})

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	var avgYearService service.GetAverageYearService
	var topSongService service.GetTopSongsService
	var topAlbumService service.GetTopAlbumsService
	var decadeBreakdownService service.GetDecadeBreakdownService
	var topArtistService service.GetTopArtistsService

	// API ENDPOINTS
	e.POST("/api/v1/updateCode", handler.HandleUpdateCode)
	e.GET("/api/v1/averageYear/:username", handler.StatsHandler(avgYearService))
	e.GET("/api/v1/topSongs/:username", handler.StatsHandler(topSongService))
	e.GET("/api/v1/topArtists/:username", handler.StatsHandler(topArtistService))
	e.GET("/api/v1/topAlbums/:username", handler.StatsHandler(topAlbumService))
	e.GET("/api/v1/decadeBreakdown/:username", handler.StatsHandler(decadeBreakdownService))

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
