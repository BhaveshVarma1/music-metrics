package main

import (
	"github.com/labstack/echo/v4"
	"music-metrics-back/handler"
)

func main() {

	e := echo.New()

	e.HTTPErrorHandler = handler.HandleNotFound
	e.Static("/", "../music-metrics-front/public")
	e.GET("/stats", handler.HandleStatic)
	e.GET("/about", handler.HandleStatic)
	e.GET("/contact", handler.HandleStatic)
	e.GET("/account", handler.HandleStatic)
	e.GET("/stats/:file", handler.HandleStatic)
	e.GET("/about/:file", handler.HandleStatic)
	e.GET("/contact/:file", handler.HandleStatic)
	e.GET("/account/:file", handler.HandleStatic)

	e.Logger.Fatal(e.Start(":3001"))

}
