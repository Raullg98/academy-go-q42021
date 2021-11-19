package handlers

import (
	"net/http"

	"go_project/models"

	"github.com/labstack/echo/v4"
)

// GetMovieById - Handles the request for getting a single movie by their id
func GetMovieById(c echo.Context) error {
	id := c.Param("id")
	movie, err := models.GetMovieById(id)
	if err != nil {
		response := map[string]string{"error": err.Error()}
		return c.JSON(http.StatusNotFound, response)
	}
	response := map[string]models.Movie{"movie": *movie}
	return c.JSON(http.StatusOK, response)
}

// GetMovieById - Handles the request for getting all the movies
func GetAllMovies(c echo.Context) error {
	movies, err := models.GetAllMovies()
	if err != nil {
		response := map[string]string{"error": err.Error()}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, map[string][]models.Movie{
		"movies": movies,
	})
}
