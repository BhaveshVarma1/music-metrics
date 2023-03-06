package handler

import (
	"github.com/labstack/echo/v4"
)

func HandleNotFound(err error, c echo.Context) {
	c.Response().WriteHeader(404)
	err = c.File("web/HTML/404.html")
}
