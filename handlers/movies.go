package handlers

import (
	"go_project/models"

	"github.com/labstack/echo/v4"
)

func GetMovieById(c echo.Context) error {
	id := c.Param("id")
	movie, err := models.GetMovieById(id)
	if err != nil {
		response := map[string]string{"error": err.Error()}
		return c.JSON(404, response)
	}
	response := map[string]models.Movie{"movie": *movie}
	return c.JSON(200, response)
}

func GetAllMovies(c echo.Context) error {
	movies, err := models.GetAllMovies()
	if err != nil {
		response := map[string]string{"error": err.Error()}
		return c.JSON(404, response)
	}

	return c.JSON(200, map[string][]models.Movie{
		"movies": movies,
	})
}
