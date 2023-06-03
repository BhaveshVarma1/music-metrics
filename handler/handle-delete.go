package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"music-metrics/da"
	"music-metrics/model"
	"music-metrics/service"
	"time"
)

func HandleDelete(c echo.Context) error {

	// Begin database transaction
	tx, db, err := da.BeginTX()
	if err != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
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
		Action:    "delete",
		IP:        c.RealIP(),
	}) != nil {
		fmt.Println("Error creating log in load handler: ", err)
	}

	// Commit changes and close the connection
	if da.CommitAndClose(tx, db, true) != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
	}

	// Call delete service
	service.Delete(username)

	return c.JSON(200, model.GenericResponse{Success: true, Message: "Successfully deleted user"})
}
