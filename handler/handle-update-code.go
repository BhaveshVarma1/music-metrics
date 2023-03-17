package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"music-metrics-back/model"
	"music-metrics-back/service"
)

func HandleUpdateCode(c echo.Context) error {

	var req model.UpdateCodeRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	resp := service.UpdateCode(req.Code)

	fmt.Println("What is up HANDLER. Resp.success: ", resp.Success)

	if resp.Success {
		fmt.Println("RETURNING 200")
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
