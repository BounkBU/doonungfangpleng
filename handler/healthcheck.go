package handler

import (
	"net/http"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

// GetServerInformation godoc
// @summary Get server information
// @tags healthcheck
// @id GetServerInformation
// @response 200 {object} model.ServerInformation "OK"
// @router / [get]
func (h *HealthcheckHandler) GetServerInformation(c *gin.Context) {
	response := model.ServerInformation{
		Project:      "Doonung-Fangpleng",
		Team:         "BounkBU",
		Collaborator: []string{"Thanathip Suwannakhot", "Pakapop Cheunchomsirakul", "Chatchawarin Chatchavalwong"},
	}
	c.JSON(http.StatusOK, response)
}

// GetHealthcheck godoc
// @summary Healthcheck
// @tags healthcheck
// @id GetHealthcheck
// @response 200 {object} model.MessageResponse "OK"
// @router /health [get]
func (h *HealthcheckHandler) GetHealthcheck(c *gin.Context) {
	message := "Hello from BounkBU"
	c.JSON(http.StatusOK, messageResponse(message))
}
