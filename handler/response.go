package handler

import (
	"errors"

	model "github.com/BounkBU/doonungfangpleng/models"
)

var ErrInvalidRequestData error = errors.New("invalid request data")

func messageResponse(message string) model.MessageResponse {
	return model.MessageResponse{
		Message: message,
	}
}

func errorResponse(err error) model.ErrorResponse {
	return model.ErrorResponse{
		Error: err.Error(),
	}
}
