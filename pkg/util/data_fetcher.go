package util

import (
	"io"
	"net/http"
)

func NewHTTPGetRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
