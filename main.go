package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-metrics-back/handler"
	"os"
	"path/filepath"
)

func main() {

	e := echo.New()

	fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))

	// todo: change this to NOT allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST("/updateCode", handler.HandleUpdateCode)

	e.Logger.Fatal(e.Start(":3001"))

}
