package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"music-metrics/da"
	"music-metrics/model"
	"music-metrics/service"
	"time"
)

func HandleLoad(c echo.Context) error {

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
	err = da.CreateLog(tx, &model.LogBean{
		Username:  username,
		Timestamp: time.Now().UnixMilli(),
		Action:    "load",
		IP:        c.RealIP(),
	})
	if err != nil {
		fmt.Println("Error creating log in load handler: ", err)
	}

	// Commit changes and close the connection
	if da.CommitAndClose(tx, db, true) != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
	}

	// Decode request body
	var req []model.ExtendedStreamingObject
	err = json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// Call load service asynchronously
	go service.Load(req, username)

	return c.JSON(200, model.GenericResponse{Success: true, Message: "Load request received"})
}
