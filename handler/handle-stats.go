package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"music-metrics/dal"
	"music-metrics/model"
	"music-metrics/service"
)

func StatsHandler(s service.StatsService) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("StatsHandler called")

		tx, db, err := dal.BeginTX()
		if err != nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
		}

		username := c.Param("username")
		token := c.Request().Header.Get("Authorization")
		authtoken, err := dal.RetrieveAuthToken(tx, token)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
			}
			return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
		}
		if authtoken.Username != username {
			if dal.CommitAndClose(tx, db, false) != nil {
				return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
			}
			return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
		}

		if dal.CommitAndClose(tx, db, true) != nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
		}

		result := s.ExecuteService(username)

		if result == nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
		}

		return c.JSON(200, result)
	}
}
