package main

import (
	"github.com/labstack/echo/v4"
	"music-metrics-back/handler"
)

func main() {

	e := echo.New()

	e.HTTPErrorHandler = handler.HandleNotFound
	e.Static("/", "../music-metrics-front/public")

	e.Logger.Fatal(e.Start(":3001"))

}
