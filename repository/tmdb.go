package repository

import (
	"encoding/json"
	"fmt"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/pkg/util"
)

type tmdbRepository struct {
	ApiKey string
}

type TmdbRepository interface {
	QueryPopularMovies(page int) ([]model.PopularTmdbResponse, error)
	QueryMovieDetails(movieId int) (model.TmdbMovieDetailRequest, error)
	QueryPopularSeries(page int) ([]model.PopularTmdbSerieRequest, error)
	QuerySerieDetails(serieId int) (model.TmdbSerieDetailRequest, error)
}

func NewTmdbRepository(apiKey string) *tmdbRepository {
	return &tmdbRepository{
		ApiKey: apiKey,
	}
}

const TMDB_URL = "https://api.themoviedb.org/3"

func (t *tmdbRepository) QueryPopularMovies(page int) (tmdbMovies []model.PopularTmdbResponse, err error) {
	popularMovieUrl := fmt.Sprintf("%s/movie/popular?api_key=%s&language=en-US&page=%d", TMDB_URL, t.ApiKey, page)
	responseData, err := util.NewHTTPGetRequest(popularMovieUrl)
	if err != nil {
		return
	}

	data := model.PopularTmdbMovie{}
	if err = json.Unmarshal(responseData, &data); err != nil {
		return
	}

	tmdbMovies = data.Results

	return tmdbMovies, nil
}

func (t *tmdbRepository) QueryMovieDetails(movieId int) (tmdbMovie model.TmdbMovieDetailRequest, err error) {
	movieDetailUrl := fmt.Sprintf("%s/movie/%d?api_key=%s&language=en-US", TMDB_URL, movieId, t.ApiKey)
	responseData, err := util.NewHTTPGetRequest(movieDetailUrl)
	if err != nil {
		return
	}

	if err = json.Unmarshal(responseData, &tmdbMovie); err != nil {
		return
	}

	return tmdbMovie, nil
}

func (t *tmdbRepository) QueryPopularSeries(page int) (tmdbSeries []model.PopularTmdbSerieRequest, err error) {
	popularSerieUrl := fmt.Sprintf("%s/tv/popular?api_key=%s&language=en-US&page=%d", TMDB_URL, t.ApiKey, page)
	responseData, err := util.NewHTTPGetRequest(popularSerieUrl)
	if err != nil {
		return
	}

	data := model.PopularTmdbSerie{}
	if err = json.Unmarshal(responseData, &data); err != nil {
		return
	}

	tmdbSeries = data.Results

	return tmdbSeries, nil
}

func (t *tmdbRepository) QuerySerieDetails(serieId int) (tmdbSerie model.TmdbSerieDetailRequest, err error) {
	serieDetailUrl := fmt.Sprintf("%s/tv/%d?api_key=%s&language=en-US", TMDB_URL, serieId, t.ApiKey)
	responseData, err := util.NewHTTPGetRequest(serieDetailUrl)
	if err != nil {
		return
	}

	if err = json.Unmarshal(responseData, &tmdbSerie); err != nil {
		return
	}

	return tmdbSerie, nil
}
