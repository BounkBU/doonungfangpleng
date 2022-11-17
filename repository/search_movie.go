package repository

import (
	model "github.com/BounkBU/doonungfangpleng/models"
	"github.com/jmoiron/sqlx"
)

type SearchMovieRepository struct {
	db *sqlx.DB
}

type SearchMovieRepositoryInterface interface {
	InsertSearchMovie(searchMovie *model.SearchMovie) error
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
