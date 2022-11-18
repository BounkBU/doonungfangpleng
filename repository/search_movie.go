package repository

import (
	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/jmoiron/sqlx"
)

type searchMovieRepository struct {
	db *sqlx.DB
}

type SearchMovieRepository interface {
	SelectSearchMovieByTMDBMovieId(tmdbMovieId int) (searchMovie model.SearchMovie, err error)
	InsertSearchMovie(searchMovie model.CreateSearchMovieRequest) error
	UpdateSearchMovieSearchAmount(tmdbMovieId int) error
	SelectPopularSearchMovies() (searchMovies []model.SearchMovie, err error)
}

func NewSearchMovieRepository(db *sqlx.DB) *searchMovieRepository {
	return &searchMovieRepository{
		db: db,
	}
}

func (s *searchMovieRepository) SelectSearchMovieByTMDBMovieId(tmdbMovieId int) (searchMovie model.SearchMovie, err error) {
	err = s.db.Get(&searchMovie, "SELECT * from search_movies WHERE tmdb_movie_id=?", tmdbMovieId)
	if err != nil {
		return
	}

	return searchMovie, nil
}

func (s *searchMovieRepository) InsertSearchMovie(searchMovie model.CreateSearchMovieRequest) error {
	_, err := s.db.Query(`
		INSERT INTO search_movies (tmdb_movie_id, title, overview, genres, image_path, release_date, rating)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		searchMovie.TMDBMovieId,
		searchMovie.Title,
		searchMovie.Overview,
		searchMovie.Genres,
		searchMovie.ImagePath,
		searchMovie.ReleaseDate,
		searchMovie.Rating,
	)
	return err
}

func (s *searchMovieRepository) UpdateSearchMovieSearchAmount(tmdbMovieId int) error {
	_, err := s.db.Query("UPDATE search_movies SET search_amount = search_amount + 1 WHERE tmdb_movie_id=?", tmdbMovieId)
	return err
}

func (s *searchMovieRepository) SelectPopularSearchMovies() ([]model.SearchMovie, error) {
	var searchMovies []model.SearchMovie
	err := s.db.Select(&searchMovies, `
		SELECT * FROM search_movies
		ORDER BY search_amount DESC
		LIMIT 10
	`)
	if err != nil {
		return searchMovies, err
	}

	return searchMovies, nil
}
