package handler

import (
	"github.com/labstack/echo/v4"
	"music-metrics-back/dal"
	"music-metrics-back/model"
	"music-metrics-back/service"
)

func HandleAverageYear(c echo.Context) error {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error2"})
	}

	username := c.Param("username")
	token := c.Request().Header.Get("Authorization")
	authtoken, err := dal.RetrieveAuthToken(tx, token)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error3"})
		}
		return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
	}
	if authtoken.Username != username {
		if dal.CommitAndClose(tx, db, false) != nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error4"})
		}
		return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
	}

	if dal.CommitAndClose(tx, db, true) != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error5"})
	}

	result := service.GetAverageYear(username)

	if result == -1 {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error6"})
	}

	return c.JSON(200, model.AverageYearResponse{Success: true, AverageYear: result})
}
