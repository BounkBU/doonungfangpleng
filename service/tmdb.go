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
	GetPopularMovieFromTmdb(page int) ([]model.PopularTmdbResponse, error)
	GetMovieDetailFromTmdb(movieId int) (model.TmdbMovieDetailResponse, error)
	GetPopularSeriesFromTmdb(page int) ([]model.PopularTmdbResponse, error)
	GetSerieDetailFromTmdb(movieId int) (model.TmdbSerieDetailResponse, error)
}

func NewTmdbService(tmdbRepository repository.TmdbRepository) *tmdbService {
	return &tmdbService{
		tmdbRepository: tmdbRepository,
	}
}

func (t *tmdbService) GetPopularMovieFromTmdb(page int) (tmdbMovies []model.PopularTmdbResponse, err error) {
	log.Info("Start getting popular movies from tmdb api")
	defer log.Info("End getting popular movies from tmdb api")

	movies, err := t.tmdbRepository.QueryPopularMovies(page)
	if err != nil {
		log.Errorf("Error query popular movies from tmdb api: %s", err.Error())
		return movies, err
	}

	for _, movie := range movies {
		tmdbMovie := model.PopularTmdbResponse{
			ID:         movie.ID,
			Title:      movie.Title,
			PosterPath: fmt.Sprintf("https://image.tmdb.org/t/p/w300_and_h450_bestv2%s", movie.PosterPath),
		}
		tmdbMovies = append(tmdbMovies, tmdbMovie)
	}

	return tmdbMovies, err
}

func (t *tmdbService) GetMovieDetailFromTmdb(movieId int) (tmdbMovie model.TmdbMovieDetailResponse, err error) {
	log.Info("Start getting movie detail from tmdb api")
	defer log.Info("End getting movie detail from tmdb api")

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

func (t *tmdbService) GetPopularSeriesFromTmdb(page int) (tmdbSeries []model.PopularTmdbResponse, err error) {
	log.Info("Start getting movie detail from tmdb api")
	defer log.Info("End getting movie detail from tmdb api")

	series, err := t.tmdbRepository.QueryPopularSeries(page)
	if err != nil {
		log.Errorf("Error query popular series from tmdb api: %s", err.Error())
		return
	}

	for _, movie := range series {
		tmdbSerie := model.PopularTmdbResponse{
			ID:         movie.ID,
			Title:      movie.Name,
			PosterPath: fmt.Sprintf("https://image.tmdb.org/t/p/w300_and_h450_bestv2%s", movie.PosterPath),
		}
		tmdbSeries = append(tmdbSeries, tmdbSerie)
	}

	return tmdbSeries, err
}

func (t *tmdbService) GetSerieDetailFromTmdb(movieId int) (tmdbSerie model.TmdbSerieDetailResponse, err error) {
	log.Info("Start getting serie detail from tmdb api")
	defer log.Info("End getting serie detail from tmdb api")

	serie, err := t.tmdbRepository.QuerySerieDetails(movieId)
	if err != nil {
		log.Errorf("Error query movie detail with id: %d: %s", movieId, err.Error())
		return
	}

	var genres []string
	for _, genre := range serie.Genres {
		genres = append(genres, genre.Name)
	}

	var seasons []model.Season
	for _, season := range serie.Seasons {
		season := model.Season{
			Name:         season.Name,
			AirDate:      season.AirDate,
			EpisodeCount: season.EpisodeCount,
			PosterPath:   fmt.Sprintf("https://image.tmdb.org/t/p/w130_and_h195_bestv2%s", serie.PosterPath),
		}
		seasons = append(seasons, season)
	}

	tmdbSerie = model.TmdbSerieDetailResponse{
		ID:              serie.ID,
		Title:           serie.Name,
		BackdropPath:    fmt.Sprintf("https://www.themoviedb.org/t/p/w1920_and_h800_multi_faces%s", serie.BackdropPath),
		Genres:          strings.Join(genres, ", "),
		Overview:        serie.Overview,
		ImagePath:       fmt.Sprintf("https://image.tmdb.org/t/p/w300_and_h450_bestv2%s", serie.PosterPath),
		NumberOfSeasons: serie.NumberOfSeasons,
		Seasons:         seasons,
		Rating:          serie.VoteAverage,
	}

	return tmdbSerie, err
}
