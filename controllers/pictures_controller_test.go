package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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

// TestGetPictureOfTheDay - Tests GetPictureOfTheDay
func TestGetPictureOfTheDay(t *testing.T) {

	testCases := []struct {
		name     string
		status   int
		picture  *models.Picture
		err      error
		date     string
		response string
	}{
		{
			name:     "find picture",
			status:   200,
			picture:  picture,
			err:      nil,
			date:     "2021-12-01",
			response: "{\"picture\":{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"}}\n",
		},
		{
			name:     "invalid date format",
			status:   400,
			picture:  picture,
			err:      nil,
			date:     "123-1234-12",
			response: "{\"error\":\"The date param must be in YYYY-MM-DD format\"}\n",
		},
		{
			name:     "service error",
			status:   500,
			picture:  nil,
			err:      errors.New("Service Error"),
			date:     "2021-12-01",
			response: "{\"error\":\"Service Error\"}\n",
		},
		{
			name:     "not found",
			status:   404,
			picture:  nil,
			err:      nil,
			date:     "2021-12-01",
			response: "{\"error\":\"Not found\"}\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.IPictureService{}
			mock.On("GetPictureFromDate", tc.date).Return(tc.picture, tc.err)

			q := make(url.Values)
			if tc.date != "" {
				q.Set("date", tc.date)
			}

			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/picture")
			h := NewPictureController(&mock)

			h.GetPictureOfTheDay(c)

			assert.Equal(t, tc.status, rec.Code)
			assert.Equal(t, tc.response, rec.Body.String())
		})

	}
}

// TestGetPictures - Tests GetPicture
func TestGetPictures(t *testing.T) {

	testCases := []struct {
		name           string
		status         int
		pictures       []models.Picture
		err            error
		criteria       string
		items          string
		itemsPerWorker string
		response       string
	}{
		{
			name:           "find even pictures",
			status:         200,
			pictures:       pictures,
			err:            nil,
			criteria:       "Even",
			items:          "10",
			itemsPerWorker: "2",
			response:       "{\"pictures\":[{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"},{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"},{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"}]}\n",
		},
		{
			name:           "find odd pictures",
			status:         200,
			pictures:       pictures,
			err:            nil,
			criteria:       "Odd",
			items:          "10",
			itemsPerWorker: "2",
			response:       "{\"pictures\":[{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"},{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"},{\"date\":\"2021-12-01\",\"copyright\":\"2021\",\"title\":\"The Extraordinary Spiral in LL Pegasi\",\"explanation\":\"What created the strange spiral structure on the upper left?\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2111/LLPegasi_HubbleLodge_1926.jpg\",\"media_type\":\"image\"}]}\n",
		},
		{
			name:           "service error",
			status:         500,
			pictures:       nil,
			err:            errors.New("Service error"),
			criteria:       "Odd",
			items:          "10",
			itemsPerWorker: "2",
			response:       "{\"error\":\"Service error\"}\n",
		},
		{
			name:           "invalid type param",
			status:         400,
			pictures:       nil,
			err:            nil,
			criteria:       "Not valid",
			items:          "10",
			itemsPerWorker: "2",
			response:       "{\"error\":\"The param type must be Even or Odd\"}\n",
		},
		{
			name:           "invalid items param",
			status:         400,
			pictures:       nil,
			err:            nil,
			criteria:       "Even",
			items:          "Not valid",
			itemsPerWorker: "2",
			response:       "{\"error\":\"The param items must be a valid integer\"}\n",
		},
		{
			name:           "invalid itemsPerWorker param",
			status:         400,
			pictures:       nil,
			err:            nil,
			criteria:       "Even",
			items:          "10",
			itemsPerWorker: "Not valid",
			response:       "{\"error\":\"The param itemsPerWorker must be a valid integer\"}\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			serviceMock := mocks.IPictureService{}
			serviceMock.On("GetPictures", tc.criteria, mock.Anything, mock.Anything).Return(tc.pictures, tc.err)

			q := make(url.Values)
			q.Set("type", tc.criteria)
			q.Set("items", tc.items)
			q.Set("itemsPerWorker", tc.itemsPerWorker)

			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/pictures")
			h := NewPictureController(&serviceMock)

			h.GetPictures(c)

			assert.Equal(t, tc.status, rec.Code)
			assert.Equal(t, tc.response, rec.Body.String())
		})

	}
}
