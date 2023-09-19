package handler

import (
	"github.com/labstack/echo/v4"
	"music-metrics/service"
)

func HandleAnalytics(c echo.Context) error {

	return c.JSON(200, service.Analytics())
}
