package handler

import (
	"net/http"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/service"
	"github.com/gin-gonic/gin"
)

type SearchMovieHandler struct {
	searchMovieService *service.SearchMovieService
}

func NewSearchMovieHandler(searchMovieService *service.SearchMovieService) *SearchMovieHandler {
	return &SearchMovieHandler{
		searchMovieService: searchMovieService,
	}
}

func (h *SearchMovieHandler) CreateSearchMovieHandler(c *gin.Context) {
	var req model.CreateSearchMovieRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := h.searchMovieService.CreateOrUpdateSearchMovie(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := "Create search movie successfully"
	c.JSON(http.StatusOK, messageResponse(response))
}
