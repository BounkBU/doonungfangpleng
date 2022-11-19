package service

import (
	"fmt"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/repository"
)

type tmdbService struct {
	TmdbRepository repository.TmdbRepository
}

type TmdbService interface {
	GetPopularMovieFromTmdb(page int) ([]model.TmdbMovie, error)
}

func NewTmdbService(tmdbRepository repository.TmdbRepository) *tmdbService {
	return &tmdbService{
		TmdbRepository: tmdbRepository,
	}
}

func (t *tmdbService) GetPopularMovieFromTmdb(page int) ([]model.TmdbMovie, error) {
	movies, err := t.TmdbRepository.QueryPopularMovies(page)

	tmdbMovies := []model.TmdbMovie{}
	for _, movie := range movies {
		tmdbMovie := model.TmdbMovie{
			ID:         movie.ID,
			Title:      movie.Title,
			PosterPath: fmt.Sprintf("https://image.tmdb.org/t/p/w300_and_h450_bestv2%s", movie.PosterPath),
		}
		tmdbMovies = append(tmdbMovies, tmdbMovie)
	}

	return tmdbMovies, err
}
