package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"music-metrics/model"
	"music-metrics/service"
)

func HandleUpdateCode(c echo.Context) error {

	// Decode request
	var req model.UpdateCodeRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// Call service
	resp := service.UpdateCode(req.Code, c.RealIP())
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
