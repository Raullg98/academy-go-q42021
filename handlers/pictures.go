package handlers

import (
	"net/http"
	"time"

	"go_project/models"
	"go_project/util"

	"github.com/labstack/echo/v4"
)

// GetMovieById - Handles the request for getting a single picture by their id
func GetPictureOfTheDay(c echo.Context) error {
	date := c.QueryParam("date")

	if !util.ValidateDateFormat(date) {
		response := map[string]string{"error": "The date param must be in YYYY-MM-DD format"}
		return c.JSON(http.StatusNotFound, response)
	}

	if date == "" {
		date = util.GetDateFormmated(time.Now())
	}

	picture, err := models.GetPictureFromDate(date)

	if err != nil {
		response := map[string]string{"error": err.Error()}
		return c.JSON(http.StatusNotFound, response)
	}
	response := map[string]models.Picture{"picture": *picture}

	err = util.AppendToCsvFile("./data/pictures.csv", picture.Values())

	if err != nil {
		response := map[string]string{"error": err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}
