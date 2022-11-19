package service

import (
	"fmt"
	"strings"

	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/BounkBU/doonungfangpleng/repository"
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

func (t *tmdbService) GetPopularMovieFromTmdb(page int) ([]model.PopularTmdbMovieResponse, error) {
	movies, err := t.tmdbRepository.QueryPopularMovies(page)

	tmdbMovies := []model.PopularTmdbMovieResponse{}
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

func (t *tmdbService) GetMovieDetailFromTmdb(movieId int) (model.TmdbMovieDetailResponse, error) {
	movie, err := t.tmdbRepository.QueryMovieDetails(movieId)

	var genres []string

	for _, genre := range movie.Genres {
		genres = append(genres, genre.Name)
	}

	tmdbMovie := model.TmdbMovieDetailResponse{
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
