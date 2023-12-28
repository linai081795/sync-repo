package utils

import (
	"github.com/go-resty/resty/v2"
)

var requestRequest *resty.Request

func Request() *resty.Request {
	if requestRequest == nil {
		request := InitRequest()

		requestRequest = request
	}
	return requestRequest
}

func InitRequest() *resty.Request {
	var client = resty.New()

	return client.R()
}
