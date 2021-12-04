package router

import (
	"github.com/labstack/echo/v4"

	"go_project/controllers"
	"go_project/repositories"
	"go_project/services"
)

func Run() {
	e := echo.New()

	pictureRepo := repositories.NewPictureRepository()
	pictureService := services.NewPictureService(pictureRepo)
	pictureController := controllers.NewPictureController(pictureService)

	e.GET("/picture", pictureController.GetPictureOfTheDay)
	e.GET("/pictures", pictureController.GetPictures)
	e.Logger.Fatal(e.Start("localhost:5000"))
}
