package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics-back/handler"
	"net/http"
)

func main() {

	e := echo.New()

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../music-metrics-front/build/index.html")
	})
	e.POST("/updateCode", handler.HandleUpdateCode)
	e.GET("/averageYear/:username", handler.HandleAverageYear)
	e.GET("/songCounts/:username", handler.HandleSongCounts)

	e.Logger.Fatal(e.Start(":3000"))

}
