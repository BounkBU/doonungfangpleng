package handler

import model "github.com/BounkBU/doonungfangpleng/models"

func messageResponse(message string) model.MessageResponse {
	return model.MessageResponse{
		Message: message,
	}
}
