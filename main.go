package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"music-metrics/handler"
	"music-metrics/service"
	"os"
)

func main() {

	buildPath := "public/build"

	e := echo.New()

	// Get the port from the command line
	port := os.Args[1]
	if port == "" || port[0] != ':' {
		fmt.Println("Please provide a port number (format-> :8080).")
		return
	}

	// Handle WebSocket connections
	e.GET("/ws", handler.HandleWebsocket)

	var allStatsService service.AllStatsService

	// API ENDPOINTS
	e.PUT("/code", handler.HandleUpdateCode)
	e.GET("/stats/:username/:range", handler.StatsHandler(allStatsService))
	e.POST("/data/:username", handler.HandleLoad)
	e.DELETE("/data/:username", handler.HandleDelete)
	e.GET("/analytics", handler.HandleAnalytics)

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

	e.Logger.Fatal(e.Start(port))
}
