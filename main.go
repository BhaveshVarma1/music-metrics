package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics/handler"
)

func main() {

	buildPath := "public/build"

	e := echo.New()

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// API ENDPOINTS
	e.POST("/api/v1/updateCode", handler.HandleUpdateCode)
	e.GET("/api/v1/averageYear/:username", handler.HandleAverageYear)
	e.GET("/api/v1/songCounts/:username", handler.HandleSongCounts)
	e.GET("/api/v1/topAlbums/:username", handler.HandleTopAlbums)

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
