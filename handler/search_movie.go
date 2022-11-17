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

// CreateSearchMovieHandler godoc
// @summary Create new search movie if does not exist, otherwise update search amount
// @tags movies
// @id CreateSearchMovie
// @accept json
// @produce json
// @param SearchMovie body model.CreateSearchMovieRequest true "Search movie data to be created"
// @response 200 {object} model.MessageResponse "OK"
// @response 400 {object} model.ErrorResponse "Bad Request"
// @response 500 {object} model.ErrorResponse "Internal Server Error"
// @router /movies [post]
func (h *SearchMovieHandler) CreateSearchMovieHandler(c *gin.Context) {
	var req model.CreateSearchMovieRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
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

// GetPopularSearchMoviesHandler godoc
// @summary Get search movies sorted by popularity
// @tags movies
// @id GetPopularSearchMovies
// @produce json
// @response 200 {array} model.SearchMovie "OK"
// @response 500 {object} model.ErrorResponse "Internal Server Error"
// @router /movies [get]
func (h *SearchMovieHandler) GetPopularSearchMoviesHandler(c *gin.Context) {
	searchMovies, err := h.searchMovieService.GetPopularSearchMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, searchMovies)
}
