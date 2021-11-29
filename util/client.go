package util

import (
	"os"

	"gopkg.in/resty.v1"
)

func CreateClient() *resty.Request {
	return resty.R().SetQueryParam("api_key", os.Getenv("API_KEY"))
}
