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

	// API ENDPOINTS
	e.PUT("/code", handler.HandleUpdateCode)
	e.GET("/stats/:username/:range", handler.StatsHandler(allStatsService))
	e.POST("/data/:username", handler.HandleLoad)
	e.DELETE("/data/:username", handler.HandleDelete)

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
