package model

type PopularTmdbMovie struct {
	Page         int64                 `json:"page"`
	Results      []PopularTmdbResponse `json:"results"`
	TotalPages   int64                 `json:"total_pages"`
	TotalResults int64                 `json:"total_results"`
}

type TmdbMovieDetailRequest struct {
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

type PopularTmdbSerie struct {
	Page         int64                     `json:"page"`
	Results      []PopularTmdbSerieRequest `json:"results"`
	TotalPages   int64                     `json:"total_pages"`
	TotalResults int64                     `json:"total_results"`
}

type PopularTmdbSerieRequest struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
}

type PopularTmdbResponse struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
}

type TmdbSerieDetailRequest struct {
	ID              int64    `json:"id"`
	BackdropPath    string   `json:"backdrop_path"`
	Genres          []Genre  `json:"genres"`
	Name            string   `json:"name"`
	NumberOfSeasons int64    `json:"number_of_seasons"`
	Overview        string   `json:"overview"`
	PosterPath      string   `json:"poster_path"`
	Seasons         []Season `json:"seasons"`
	VoteAverage     float64  `json:"vote_average"`
}

type TmdbSerieDetailResponse struct {
	ID              int64    `json:"tmdb_id"`
	Title           string   `json:"title"`
	BackdropPath    string   `json:"backdrop_path"`
	Genres          string   `json:"genres"`
	Overview        string   `json:"overview"`
	ImagePath       string   `json:"image_path"`
	NumberOfSeasons int64    `json:"number_of_seasons"`
	Seasons         []Season `json:"seasons"`
	Rating          float64  `json:"rating"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Season struct {
	AirDate      string `json:"air_date"`
	Name         string `json:"name"`
	EpisodeCount int64  `json:"episode_count"`
	PosterPath   string `json:"poster_path"`
}
