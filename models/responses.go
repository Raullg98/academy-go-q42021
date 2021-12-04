package models

import (
	"errors"
)

type PictureResponse struct {
	Picture Picture `json:"picture"`
}

type PicturesResponse struct {
	Pictures []Picture `json:"pictures"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type NasaErrorResponse struct {
	Error NasaApiErrorResponse `json:"error"`
	GeneralErrorResponse
}

type NasaApiErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type GeneralErrorResponse struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	ServiceVersion string `json:"service_version"`
}

// GetError - Returns the error message
func (err *NasaErrorResponse) GetError() error {
	if err.Msg != "" {
		return errors.New(err.Msg)
	}
	return errors.New(err.Error.Msg)
}
