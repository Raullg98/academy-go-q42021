package controllers

import (
	"net/http"
	"strconv"

	"go_project/interfaces"
	"go_project/models"
	"go_project/responses"
	"go_project/util"

	"github.com/labstack/echo/v4"
)

type PictureController struct {
	service interfaces.IPictureService
}

// NewPictureController - Creates a new picture controller
func NewPictureController(ps interfaces.IPictureService) *PictureController {
	return &PictureController{ps}
}

// GetPictureOfTheDay - Handles the request for getting a single picture by their date
func (pc *PictureController) GetPictureOfTheDay(c echo.Context) error {
	date := c.QueryParam("date")

	if date != "" && !util.ValidateDateFormat(date) {
		return responses.BadRequestError(c, "The date param must be in YYYY-MM-DD format")
	}

	picture, err := pc.service.GetPictureFromDate(date)

	if err != nil {
		return responses.InternalServerError(c, err)
	}

	if picture == nil {
		return responses.NotFoundError(c)
	}

	response := models.PictureResponse{Picture: *picture}
	return c.JSON(http.StatusOK, response)
}

// GetPictures - Returns a list of picture based on criteria, items and itemsPerWorker
func (pc *PictureController) GetPictures(c echo.Context) error {

	criteria := c.QueryParam("type")

	if criteria != "Even" && criteria != "Odd" {
		return responses.BadRequestError(c, "The param type must be Even or Odd")
	}

	items, err := strconv.Atoi(c.QueryParam("items"))
	if err != nil {
		return responses.BadRequestError(c, "The param items must be a valid integer")
	}

	itemsPerWorker, err := strconv.Atoi(c.QueryParam("itemsPerWorker"))
	if err != nil {
		return responses.BadRequestError(c, "The param itemsPerWorker must be a valid integer")
	}

	pictures, err := pc.service.GetPictures(criteria, items, itemsPerWorker)
	if err != nil {
		return responses.InternalServerError(c, err)
	}
	response := models.PicturesResponse{Pictures: pictures}

	return c.JSON(http.StatusOK, response)
}
