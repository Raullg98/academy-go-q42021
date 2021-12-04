package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "go_project/mocks/interfaces"
	"go_project/models"
)

var (
	picture = &models.Picture{
		Date:        "2021-12-01",
		Copyright:   "2021",
		Title:       "The Extraordinary Spiral in LL Pegasi",
		Explanation: "What created the strange spiral structure on the upper left?",
		Url:         "https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg",
		MediaType:   "image",
	}

	pictures = []models.Picture{
		{
			Date:        "2021-12-01",
			Copyright:   "2021",
			Title:       "The Extraordinary Spiral in LL Pegasi",
			Explanation: "What created the strange spiral structure on the upper left?",
			Url:         "https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg",
			MediaType:   "image",
		},
		{
			Date:        "2021-12-01",
			Copyright:   "2021",
			Title:       "The Extraordinary Spiral in LL Pegasi",
			Explanation: "What created the strange spiral structure on the upper left?",
			Url:         "https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg",
			MediaType:   "image",
		},
		{
			Date:        "2021-12-01",
			Copyright:   "2021",
			Title:       "The Extraordinary Spiral in LL Pegasi",
			Explanation: "What created the strange spiral structure on the upper left?",
			Url:         "https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg",
			MediaType:   "image",
		},
	}
)

// GetPictureOfTheDay - Handles the request for getting a single picture by their id
func TestGetPictureFromDate(t *testing.T) {

	testCases := []struct {
		name             string
		status           int
		picture          *models.Picture
		err              error
		savePictureError error
		date             string
		response         string
	}{
		{
			name:    "find picture",
			picture: picture,
			err:     nil,
			date:    "2021-12-01",
		},
		{
			name:    "repository error",
			picture: nil,
			err:     nil,
			date:    "2021-12-01",
		},
		{
			name:             "save csv error",
			picture:          picture,
			err:              nil,
			date:             "2021-12-01",
			savePictureError: errors.New("Csv Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.IPictureRepository{}
			mock.On("FindPictureByDate", tc.date).Return(tc.picture, tc.err)
			mock.On("SavePicture", tc.picture).Return(tc.savePictureError)

			service := NewPictureService(&mock)

			picture, err := service.GetPictureFromDate(tc.date)
			if tc.err != nil {
				assert.EqualError(t, tc.err, err.Error())
			} else if tc.savePictureError != nil {
				assert.EqualError(t, tc.savePictureError, err.Error())
			} else if tc.picture != nil {
				assert.Equal(t, tc.picture, picture)
			}

		})

	}
}

func TestGetPictures(t *testing.T) {

	testCases := []struct {
		name           string
		pictures       []models.Picture
		err            error
		criteria       string
		items          int
		itemsPerWorker int
	}{
		{
			name:           "find pictures",
			pictures:       pictures,
			err:            nil,
			criteria:       "Even",
			items:          10,
			itemsPerWorker: 2,
		},
		{
			name:           "repository error",
			pictures:       pictures,
			err:            nil,
			criteria:       "Even",
			items:          10,
			itemsPerWorker: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.IPictureRepository{}
			mock.On("ReadPicturesWithWorkerPool", tc.criteria, tc.items, tc.itemsPerWorker).Return(tc.pictures, tc.err)

			service := NewPictureService(&mock)

			pictures, err := service.GetPictures(tc.criteria, tc.items, tc.itemsPerWorker)
			if tc.err != nil {
				assert.EqualError(t, tc.err, err.Error())
			}

			assert.Equal(t, tc.pictures, pictures)
		})

	}
}
