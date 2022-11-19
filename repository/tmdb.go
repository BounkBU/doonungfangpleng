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
	QueryPopularMovies(page int) ([]model.TmdbMovie, error)
}

func NewTmdbRepository(apiKey string) *tmdbRepository {
	return &tmdbRepository{
		ApiKey: apiKey,
	}
}

const TMDB_URL = "https://api.themoviedb.org/3"

func (t *tmdbRepository) QueryPopularMovies(page int) (tmdbMovies []model.TmdbMovie, err error) {
	popularMovieUrl := fmt.Sprintf("%s/movie/popular?api_key=%s&language=en-US&page=%d", TMDB_URL, t.ApiKey, page)
	responseData, err := util.FetchData(popularMovieUrl)
	if err != nil {
		return
	}

	data := model.AllTmdbMovieResult{}
	if err = json.Unmarshal(responseData, &data); err != nil {
		return
	}

	tmdbMovies = data.Results

	return tmdbMovies, nil
}
