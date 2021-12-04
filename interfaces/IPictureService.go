package interfaces

import "go_project/models"

type IPictureService interface {
	GetPictureFromDate(date string) (*models.Picture, error)
	GetPictures(criteria string, items, itemsPerWorker int) ([]models.Picture, error)
}
