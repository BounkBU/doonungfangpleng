package model

type PopularTmdbMovie struct {
	Page         int64                      `json:"page"`
	Results      []PopularTmdbMovieResponse `json:"results"`
	TotalPages   int64                      `json:"total_pages"`
	TotalResults int64                      `json:"total_results"`
}

type PopularTmdbMovieResponse struct {
	ID         int64  `json:"id"`
	PosterPath string `json:"poster_path"`
	Title      string `json:"title"`
}

type TmdbMovieDetail struct {
	Genres       []Genre `json:"genres"`
	ID           int64   `json:"id"`
	Overview     string  `json:"overview"`
	BackdropPath string  `json:"backdrop_path"`
	PosterPath   string  `json:"poster_path"`
	ReleaseDate  string  `json:"release_date"`
	Title        string  `json:"title"`
	VoteAverage  float64 `json:"vote_average"`
}

type TmdbMovieDetailResponse struct {
	ID           int64   `json:"tmdb_id"`
	Title        string  `json:"title"`
	BackdropPath string  `json:"backdrop_path"`
	Genres       string  `json:"genres"`
	Overview     string  `json:"overview"`
	ImagePath    string  `json:"image_path"`
	ReleaseDate  string  `json:"release_date"`
	Rating       float64 `json:"rating"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
