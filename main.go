package main

import (
	"github.com/labstack/echo/v4"
	"music-metrics-back/handler"
)

func main() {

	e := echo.New()

	e.POST("/updateCode", handler.HandleUpdateCode)

	e.Logger.Fatal(e.Start(":3001"))

}
