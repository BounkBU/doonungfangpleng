package service

import (
	"fmt"
	"strings"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/repository"
	log "github.com/sirupsen/logrus"
)

type tmdbService struct {
	tmdbRepository repository.TmdbRepository
}

type TmdbService interface {
	GetPopularMovieFromTmdb(page int) ([]model.PopularTmdbMovieResponse, error)
	GetMovieDetailFromTmdb(movieId int) (model.TmdbMovieDetailResponse, error)
}

func NewTmdbService(tmdbRepository repository.TmdbRepository) *tmdbService {
	return &tmdbService{
		tmdbRepository: tmdbRepository,
	}
}

func (t *tmdbService) GetPopularMovieFromTmdb(page int) (tmdbMovies []model.PopularTmdbMovieResponse, err error) {
	log.Info("Start getting popular movies from tmdb api")
	defer log.Info("End getting popular movies from tmdb api")

	movies, err := t.tmdbRepository.QueryPopularMovies(page)
	if err != nil {
		log.Errorf("Error query popular movies from tmdb api: %s", err.Error())
		return movies, err
	}

	for _, movie := range movies {
		tmdbMovie := model.PopularTmdbMovieResponse{
			ID:         movie.ID,
			Title:      movie.Title,
			PosterPath: fmt.Sprintf("https://image.tmdb.org/t/p/w300_and_h450_bestv2%s", movie.PosterPath),
		}
		tmdbMovies = append(tmdbMovies, tmdbMovie)
	}

	return tmdbMovies, err
}

func (t *tmdbService) GetMovieDetailFromTmdb(movieId int) (tmdbMovie model.TmdbMovieDetailResponse, err error) {
	log.Info("Start getting popular movies from tmdb api")
	defer log.Info("End getting popular movies from tmdb api")

	movie, err := t.tmdbRepository.QueryMovieDetails(movieId)
	if err != nil {
		log.Errorf("Error query movie detail with id: %d: %s", movieId, err.Error())
		return
	}

	var genres []string
	for _, genre := range movie.Genres {
		genres = append(genres, genre.Name)
	}

	tmdbMovie = model.TmdbMovieDetailResponse{
		ID:           movie.ID,
		Overview:     movie.Overview,
		Genres:       strings.Join(genres, ", "),
		BackdropPath: fmt.Sprintf("https://www.themoviedb.org/t/p/w1920_and_h800_multi_faces%s", movie.BackdropPath),
		ImagePath:    fmt.Sprintf("https://image.tmdb.org/t/p/w300_and_h450_bestv2%s", movie.PosterPath),
		ReleaseDate:  movie.ReleaseDate,
		Title:        movie.Title,
		Rating:       movie.VoteAverage,
	}

	return tmdbMovie, err
}
