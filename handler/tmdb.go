package handler

import (
	"net/http"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/service"
	"github.com/gin-gonic/gin"
)

type tmdbHandler struct {
	TmdbService service.TmdbService
}

func NewTmdbHandler(tmdbService service.TmdbService) tmdbHandler {
	return tmdbHandler{
		TmdbService: tmdbService,
	}
}

func (h *tmdbHandler) GetAllPopularMovieFromTmdbHandler(c *gin.Context) {
	var req model.GetAllPopularMovieFromTmdbRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
		return
	}

	data, err := h.TmdbService.GetPopularMovieFromTmdb(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, data)
}
