CREATE TABLE `search_movies` (
  `id` bigserial PRIMARY KEY AUTO_INCREMENT,
  `tmdb_movie_id` varchar(255),
  `title` varchar(255),
  `overview` varchar(255),
  `genres` varchar(255),
  `image_path` varchar(255),
  `release_date` varchar(255),
  `rating` float,
  `search_amount` bigint NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `playlists` (
  `id` bigserial PRIMARY KEY AUTO_INCREMENT,
  `movie_id` varchar(255),
  `spotify_playlist_id` varchar(255),
  `name` varchar(255),
  `owner` varchar(255),
  `image_path` varchar(255),
  `visit_amount` bigint NOT NULL DEFAULT 1,
  `created_at` timestamptz NOT NULL DEFAULT (now())
);
