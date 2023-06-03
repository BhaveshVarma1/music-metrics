package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"music-metrics/da"
	"music-metrics/model"
	"music-metrics/service"
	"strconv"
	"strings"
	"time"
)

func StatsHandler(s service.StatsService) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Begin database transaction
		tx, db, err := da.BeginTX()
		if err != nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
		}

		// Parse request URI
		timeRange := strings.Split(c.Param("range"), "-")
		startTime, err := strconv.ParseInt(timeRange[0], 10, 64)
		endTime, err := strconv.ParseInt(timeRange[1], 10, 64)
		if err != nil || startTime > endTime {
			return c.JSON(400, model.GenericResponse{Success: false, Message: "Bad time range"})
		}

		// Validate authtoken
		username := c.Param("username")
		token := c.Request().Header.Get("Authorization")
		authtoken, err := da.RetrieveAuthToken(tx, token)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
			}
			return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
		}
		if authtoken.Username != username {
			if da.CommitAndClose(tx, db, false) != nil {
				return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
			}
			return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
		}

		// Log this request
		if da.CreateLog(tx, &model.LogBean{
			Username:  username,
			Timestamp: time.Now().UnixMilli(),
			Action:    "stats",
			IP:        c.RealIP(),
		}) != nil {
			fmt.Println("Error creating log in stats handler: ", err)
		}

		// Commit changes and close the connection
		if da.CommitAndClose(tx, db, true) != nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
		}

		// Call stats service
		result := s.ExecuteService(username, startTime, endTime)
		if result == nil {
			return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
		}

		return c.JSON(200, result)
	}
}
