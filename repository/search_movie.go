package repository

import (
	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/jmoiron/sqlx"
)

type SearchMovieRepository struct {
	db *sqlx.DB
}

type SearchMovieRepositoryInterface interface {
	SelectSearchMovieByTMDBMovieId(tmdbMovieId int) (searchMovie model.SearchMovie, err error)
	InsertSearchMovie(searchMovie *model.SearchMovie) error
	UpdateSearchMovieSearchAmount(tmdbMovieId int) error
}

func NewSearchMovieRepository(db *sqlx.DB) *SearchMovieRepository {
	return &SearchMovieRepository{
		db: db,
	}
}

func (s *SearchMovieRepository) SelectSearchMovieByTMDBMovieId(tmdbMovieId int) (searchMovie model.SearchMovie, err error) {
	err = s.db.Get(&searchMovie, "SELECT * from search_movies WHERE tmdb_movie_id=?", tmdbMovieId)
	if err != nil {
		return
	}

	return searchMovie, nil
}

func (s *SearchMovieRepository) InsertSearchMovie(searchMovie model.CreateSearchMovieRequest) error {
	_, err := s.db.Query(`
		INSERT INTO search_movies (tmdb_movie_id, title, overview, genres, image_path, release_date, rating, search_amount)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		searchMovie.TMDBMovieId,
		searchMovie.Title,
		searchMovie.Overview,
		searchMovie.Genres,
		searchMovie.ImagePath,
		searchMovie.ReleaseDate,
		searchMovie.Rating,
		1,
	)
	return err
}

func (s *SearchMovieRepository) UpdateSearchMovieSearchAmount(tmdbMovieId int) error {
	_, err := s.db.Query("UPDATE search_movies SET search_amount = search_amount + 1 WHERE tmdb_movie_id=?", tmdbMovieId)
	return err
}
