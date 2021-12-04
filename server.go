package main

import (
	"log"

	"go_project/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e := echo.New()
	e.GET("/picture", handlers.GetPictureOfTheDay)
	e.GET("/movies", handlers.GetAllMovies)
	e.GET("/movies/:id", handlers.GetMovieById)
	e.Logger.Fatal(e.Start(":5000"))
}
