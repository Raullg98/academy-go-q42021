package services

import (
	"time"

	"go_project/interfaces"
	"go_project/models"
)

type PictureService struct {
	interfaces.IPictureRepository
}

// NewPictureService - Creates a new picture service
func NewPictureService(pr interfaces.IPictureRepository) *PictureService {
	return &PictureService{pr}
}

// GetPictureFromDate - Returns the picture from a date
func (service *PictureService) GetPictureFromDate(date string) (*models.Picture, error) {
	if date == "" {
		date = time.Now().Local().Format("2006-01-02")
	}

	var picture, err = service.FindPictureByDate(date)
	if err != nil {
		return nil, err
	}

	if picture == nil {
		return nil, nil
	}

	err = service.SavePicture(picture)
	if err != nil {
		return nil, err
	}

	return picture, nil
}

// GetPictures - Returns a list of picture based on criteria, items, and itemsPerWorker
func (service *PictureService) GetPictures(criteria string, items, itemsPerWorker int) ([]models.Picture, error) {
	var pictures, err = service.ReadPicturesWithWorkerPool(criteria, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}

	return pictures, nil
}
