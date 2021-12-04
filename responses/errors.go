package responses

import (
	"net/http"

	"go_project/models"

	"github.com/labstack/echo/v4"
)

// BadRequestsError - Returns an error that represents a 400 status http response
func BadRequestError(c echo.Context, message string) error {
	res := models.ErrorResponse{
		Error: message,
	}
	return c.JSON(http.StatusBadRequest, res)
}

// NotFoundError - Returns an error that represents a 404 status http response
func NotFoundError(c echo.Context) error {
	res := models.ErrorResponse{
		Error: "Not found",
	}
	return c.JSON(http.StatusNotFound, res)
}

// InternalServerError - Returns an error that represents a 500 status http response
func InternalServerError(c echo.Context, err error) error {
	res := models.ErrorResponse{
		Error: err.Error(),
	}
	return c.JSON(http.StatusInternalServerError, res)
}
