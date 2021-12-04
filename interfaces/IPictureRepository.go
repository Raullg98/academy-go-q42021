package interfaces

import (
	"go_project/models"
)

type IPictureRepository interface {
	FindPictureByDate(date string) (*models.Picture, error)
	SavePicture(picture *models.Picture) error
	ReadPicturesWithWorkerPool(criteria string, items int, maxJobsPerWorker int) ([]models.Picture, error)
}
