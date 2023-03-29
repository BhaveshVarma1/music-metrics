package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics/handler"
	"net/http"
)

func main() {

	buildPath := "public/build"

	e := echo.New()

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST("/updateCode", handler.HandleUpdateCode)
	e.GET("/averageYear/:username", handler.HandleAverageYear)
	e.GET("/songCounts/:username", handler.HandleSongCounts)

	/*e.GET("/static/*", func(c echo.Context) error {
		fmt.Println("Serving static file")
		fmt.Println(buildPath + c.Request().URL.Path)
		return c.File(buildPath + c.Request().URL.Path)
	})

	e.GET("/manifest.json", func(c echo.Context) error {
		return c.File(buildPath + "/manifest.json")
	})*/

	e.GET("/*", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		http.FileServer(http.Dir(buildPath)).ServeHTTP(res, req)
		return nil
	})

	e.Logger.Fatal(e.Start(":3000"))

}
