package handler

import (
	"github.com/labstack/echo/v4"
)

func HandleNotFound(err error, c echo.Context) {
	c.Response().WriteHeader(404)
	err = c.File("../music-metrics-front/public/404.html")
}

func HandleStatic(c echo.Context) error {

	root := "../music-metrics-front/public/"

	switch c.Path() {
	case "/stats":
		return c.File(root + "stats.html")
	case "/about":
		return c.File(root + "about.html")
	case "/contact":
		return c.File(root + "contact.html")
	case "/account":
		return c.File(root + "account.html")
	}

	switch c.Param("file") {
	case "favicon.ico":
		return c.File(root + "favicon.ico")
	case "styles.css":
		return c.File(root + "styles.css")
	case "script.js":
		return c.File(root + "script.js")
	case "logo.png":
		return c.File(root + "logo.png")
	default:
		return c.File(root + "index.html")
	}
}
