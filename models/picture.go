package models

import (
	"go_project/util"
)

type Picture struct {
	Copyright   string `json:"copyright"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	Url         string `json:"hdurl"`
	MediaType   string `json:"media_type"`
}

func (p Picture) Values() []string {
	return []string{
		p.Copyright,
		p.Date,
		p.Title,
		p.Explanation,
		p.Url,
		p.MediaType,
	}
}

// GetAllPictures - Returns all pictures
func FetchPicture(date string) (*Picture, error) {
	client := util.CreateClient().
		SetError(&ErrorResponse{}).
		SetResult(&Picture{})

	resp, err := client.
		SetQueryParam("date", date).
		Get("https://api.nasa.gov/planetary/apod")

	if err != nil {
		return nil, err
	} else if resp.IsError() {
		responseError := resp.Error().(*ErrorResponse)
		return nil, responseError.GetError()
	}

	picture := resp.Result().(*Picture)
	return picture, nil
}

// GetPictureById - Returns a single picture based on their id
func GetPictureFromDate(date string) (*Picture, error) {
	var picture, err = FetchPicture(date)
	if err != nil {
		return nil, err
	}

	return picture, nil
}
