package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"music-metrics/dal"
	"music-metrics/model"
	"music-metrics/service"
)

func HandleLoad(c echo.Context) error {

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

	// DELETE THIS BLOCK
	/*bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)*/
	// DELETE THIS BLOCK

	// decode request body
	req := model.ExtendedStreamingHistory{}
	err = json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// call service asynchronously
	go service.Load(req, username)

	return c.JSON(200, model.GenericResponse{Success: true, Message: "Load request received"})
}
