package models

import (
	"errors"
)

type ErrorResponse struct {
	Error ApiErrorResponse `json:error`
	GeneralErrorResponse
}

type ApiErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type GeneralErrorResponse struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	ServiceVersion string `json:"service_version"`
}

// GetError - Returns the error message
func (err ErrorResponse) GetError() error {
	if err.Msg != "" {
		return errors.New(err.Msg)
	}
	return errors.New(err.Error.Msg)
}
