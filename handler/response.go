package handler

import model "github.com/BounkBU/doonungfangpleng/models"

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
