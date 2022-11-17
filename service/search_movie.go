package service

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/repository"
)

type SearchMovieService struct {
	searchMovieRepository *repository.SearchMovieRepository
}

type SearchMovieServiceInterface interface {
	CreateOrUpdateSearchMovie(newSearchMovie model.CreateSearchMovieRequest) error
}

func NewSearchMovieSearchService(searchMovieRepository *repository.SearchMovieRepository) *SearchMovieService {
	return &SearchMovieService{
		searchMovieRepository: searchMovieRepository,
	}
}

func (s *SearchMovieService) CreateOrUpdateSearchMovie(newSearchMovie model.CreateSearchMovieRequest) error {
	log.Info("Start creating or updating new search movie")
	defer log.Info("End creating or updating new search movie")
	searchMovie, err := s.searchMovieRepository.SelectSearchMovieByTMDBMovieId(newSearchMovie.TMDBMovieId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("Error select search movie by tmdb_movie_id: %s", err.Error())
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = s.searchMovieRepository.InsertSearchMovie(newSearchMovie)
		return err
	}

	err = s.searchMovieRepository.UpdateSearchMovieSearchAmount(searchMovie.TMDBMovieId)
	return err
}

func (s *SearchMovieService) GetPopularSearchMovies() ([]model.SearchMovie, error) {
	log.Info("Start getting new search movie")
	defer log.Info("End getting new search movie")

	searchMovies, err := s.searchMovieRepository.SelectPopularSearchMovies()
	if err != nil {
		log.Errorf("Error select popular search movies: %s", err.Error())
		return searchMovies, err
	}

	return searchMovies, nil
}
