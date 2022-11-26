package httpserver

import (
	"github.com/BounkBU/doonungfangpleng/config"
	_ "github.com/BounkBU/doonungfangpleng/docs"
	"github.com/BounkBU/doonungfangpleng/handler"
	"github.com/BounkBU/doonungfangpleng/repository"
	"github.com/BounkBU/doonungfangpleng/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	App      *gin.Engine
	Database *sqlx.DB
	Config   *config.Config
}

func NewHTTPServer(config *config.Config, db *sqlx.DB) *Server {
	gin.SetMode(config.App.GinMode)
	app := gin.Default()
	return &Server{
		App:      app,
		Database: db,
		Config:   config,
	}
}

// @title Doonung-FangPleng API
// @version 1.0
// @description The Doonung-FangPleng web API

// @contact.name BounkBU Support
// @contact.email thanathip.suw@gmail.com

// @license.name MIT License
// @license.url https://choosealicense.com/licenses/mit/

// @schemes http
// @host localhost:8888
func (server *Server) SetupRouter() {
	// Enable cors request
	server.App.Use(cors.Default())
	server.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	healthcheckHandler := handler.NewHealthcheckHandler()
	server.App.GET("/", healthcheckHandler.GetServerInformation)
	server.App.GET("/health", healthcheckHandler.GetHealthcheck)

	searchMovieRepository := repository.NewSearchMovieRepository(server.Database)
	searchMovieService := service.NewSearchMovieSearchService(searchMovieRepository)
	searchMovieHandler := handler.NewSearchMovieHandler(searchMovieService)
	server.App.GET("/movies", searchMovieHandler.GetPopularSearchMoviesHandler)
	server.App.POST("/movies", searchMovieHandler.CreateSearchMovieHandler)

	tmdbRepository := repository.NewTmdbRepository(server.Config.App.TmdbApiKey)
	tmdbService := service.NewTmdbService(tmdbRepository)
	tmdbHandler := handler.NewTmdbHandler(tmdbService)
	server.App.GET("/tmdb/movies", tmdbHandler.GetAllPopularMovieFromTmdbHandler)
	server.App.GET("/tmdb/movie/:movieId", tmdbHandler.GetMovieDetailHandler)
	server.App.GET("/tmdb/series", tmdbHandler.GetAllPopularSeriesFromTmdbHandler)
	server.App.GET("/tmdb/serie/:serieId", tmdbHandler.GetSerieDetailHandler)
}

func (server *Server) Start() {
	server.SetupRouter()

	port := server.Config.App.Port

	log.Infof("Server is starting on port : %s", port)
	server.App.Run(":" + port)
}
