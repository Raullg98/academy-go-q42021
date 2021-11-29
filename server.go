package main

import (
	"go_project/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.GetAllMovies)
	e.GET("/:id", handlers.GetMovieById)
	e.Logger.Fatal(e.Start(":5000"))
}
