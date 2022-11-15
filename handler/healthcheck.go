package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) GetHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from bounkBU"})
}
