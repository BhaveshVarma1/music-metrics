package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics-back/handler"
)

func main() {

	buildPath := "../music-metrics-front/build"

	e := echo.New()

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST("/updateCode", handler.HandleUpdateCode)
	e.GET("/averageYear/:username", handler.HandleAverageYear)
	e.GET("/songCounts/:username", handler.HandleSongCounts)

	e.GET("/static/*", func(c echo.Context) error {
		fmt.Println("Returning file at: " + buildPath + c.Request().URL.Path)
		return c.File(buildPath + c.Request().URL.Path)
	})

	e.GET("/manifest.json", func(c echo.Context) error {
		fmt.Println("Returning manifest.json")
		return c.File(buildPath + "/manifest.json")
	})

	e.GET("/favicon.ico", func(c echo.Context) error {
		fmt.Println("Returning favicon.ico")
		return c.File(buildPath + "/favicon.ico")
	})

	e.GET("/*", func(c echo.Context) error {
		fmt.Println("Returning index.html")
		return c.File(buildPath + "/index.html")
	})

	e.Logger.Fatal(e.Start(":3000"))

}
