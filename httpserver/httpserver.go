package httpserver

import (
	"github.com/BounkBU/doonungfangpleng/config"
	"github.com/BounkBU/doonungfangpleng/handler"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

func (server *Server) SetupRouter() {
	healthcheckHandler := handler.NewHealthcheckHandler()

	server.App.GET("/", healthcheckHandler.GetServerInformation)
	server.App.GET("/health", healthcheckHandler.GetHealthcheck)
}

func (server *Server) Start() {
	server.SetupRouter()
	server.App.Run(server.Config.App.ServerAddress)
}
