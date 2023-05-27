package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"music-metrics/da"
	"music-metrics/model"
	"music-metrics/service"
)

func HandleLoad(c echo.Context) error {

	tx, db, err := da.BeginTX()
	if err != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
	}

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

	if da.CommitAndClose(tx, db, true) != nil {
		return c.JSON(500, model.GenericResponse{Success: false, Message: "Internal server error"})
	}

	// decode request body
	var req []model.ExtendedStreamingObject
	err = json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// call service asynchronously
	go service.Load(req, username)

	return c.JSON(200, model.GenericResponse{Success: true, Message: "Load request received"})
}
