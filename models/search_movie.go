package model

import "time"

type SearchMovie struct {
	Id           int       `db:"id" json:"id"`
	TMDBMovieId  int       `db:"tmdb_movie_id" json:"tmdb_movie_id"`
	Title        string    `db:"title" json:"title"`
	Overview     string    `db:"overview" json:"overview"`
	Genres       string    `db:"genres" json:"genres"`
	ImagePath    string    `db:"image_path" json:"image_path"`
	ReleaseDate  string    `db:"release_date" json:"release_date"`
	Rating       float64   `db:"rating" json:"rating"`
	SearchAmount int       `db:"search_amount" json:"search_amount"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type CreateSearchMovieRequest struct {
	TMDBMovieId int     `db:"tmdb_movie_id" json:"tmdb_movie_id"`
	Title       string  `db:"title" json:"title"`
	Overview    string  `db:"overview" json:"overview"`
	Genres      string  `db:"genres" json:"genres"`
	ImagePath   string  `db:"image_path" json:"image_path"`
	ReleaseDate string  `db:"release_date" json:"release_date"`
	Rating      float64 `db:"rating" json:"rating"`
}
