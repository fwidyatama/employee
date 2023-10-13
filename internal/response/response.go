package response

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Message string `json:"message,omitempty"`
}

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data"`
	Error   *Error      `json:"error,omitempty"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	resp := Response{}
	resp.Message = "success"
	resp.Status = 200
	resp.Error = nil
	resp.Data = data

	return c.JSON(200, resp)
}

func ErrorResponse(c echo.Context, message string, statusCode int) error {

	resp := Response{}
	resp.Message = "failed"
	resp.Status = statusCode
	resp.Error = &Error{}
	resp.Error.Message = message

	return c.JSON(statusCode, resp)

}
