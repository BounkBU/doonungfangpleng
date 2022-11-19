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
	QueryPopularMovies(page int) ([]model.PopularTmdbMovieResponse, error)
	QueryMovieDetails(movieId int) (tmdbMovie model.TmdbMovieDetail, err error)
}

func NewTmdbRepository(apiKey string) *tmdbRepository {
	return &tmdbRepository{
		ApiKey: apiKey,
	}
}

const TMDB_URL = "https://api.themoviedb.org/3"

func (t *tmdbRepository) QueryPopularMovies(page int) (tmdbMovies []model.PopularTmdbMovieResponse, err error) {
	popularMovieUrl := fmt.Sprintf("%s/movie/popular?api_key=%s&language=en-US&page=%d", TMDB_URL, t.ApiKey, page)
	responseData, err := util.FetchData(popularMovieUrl)
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

func (t *tmdbRepository) QueryMovieDetails(movieId int) (tmdbMovie model.TmdbMovieDetail, err error) {
	movieDetailUrl := fmt.Sprintf("%s/movie/%d?api_key=%s&language=en-US", TMDB_URL, movieId, t.ApiKey)
	responseData, err := util.FetchData(movieDetailUrl)
	if err != nil {
		return
	}

	if err = json.Unmarshal(responseData, &tmdbMovie); err != nil {
		return
	}

	return tmdbMovie, nil
}
