package model

type AllTmdbMovieResult struct {
	Page         int64       `json:"page"`
	Results      []TmdbMovie `json:"results"`
	TotalPages   int64       `json:"total_pages"`
	TotalResults int64       `json:"total_results"`
}

type TmdbMovie struct {
	ID         int64  `json:"id"`
	PosterPath string `json:"poster_path"`
	Title      string `json:"title"`
}

type GetAllPopularMovieFromTmdbRequest struct {
	Page int `json:"page"`
}
