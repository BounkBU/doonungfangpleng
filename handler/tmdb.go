package handler

import (
	"net/http"
	"strconv"

	"github.com/BounkBU/doonungfangpleng/service"
	"github.com/gin-gonic/gin"
)

type tmdbHandler struct {
	tmdbService service.TmdbService
}

func NewTmdbHandler(tmdbService service.TmdbService) tmdbHandler {
	return tmdbHandler{
		tmdbService: tmdbService,
	}
}

func (h *tmdbHandler) GetAllPopularMovieFromTmdbHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidQueryParam))
		return
	}

	tmdbMovies, err := h.tmdbService.GetPopularMovieFromTmdb(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, tmdbMovies)
}

func (h *tmdbHandler) GetMovieDetailHandler(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("movieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidQueryParam))
		return
	}

	tmdbMovie, err := h.tmdbService.GetMovieDetailFromTmdb(movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, tmdbMovie)
}

func (h *tmdbHandler) GetAllPopularSeriesFromTmdbHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidQueryParam))
		return
	}

	tmdbSeries, err := h.tmdbService.GetPopularSeriesFromTmdb(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, tmdbSeries)
}

func (h *tmdbHandler) GetSerieDetailHandler(c *gin.Context) {
	serieId, err := strconv.Atoi(c.Param("serieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidQueryParam))
		return
	}

	tmdbMovie, err := h.tmdbService.GetSerieDetailFromTmdb(serieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, tmdbMovie)
}
