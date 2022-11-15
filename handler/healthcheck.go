package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) GetServerInformation(c *gin.Context) {
	response := gin.H{
		"project":     "Doonung-Fangpleng",
		"team":        "BounkBU",
		"colaborator": []string{"Thanathip Suwannakhot", "Pakapop Cheunchomsirakul", "Chatchawarin Chatchavalwong"},
	}
	c.JSON(http.StatusOK, response)
}

func (h *HealthcheckHandler) GetHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from bounkBU"})
}
