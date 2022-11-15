package httpserver

import (
	"github.com/BounkBU/doonungfangpleng/config"
	"github.com/BounkBU/doonungfangpleng/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin    *gin.Engine
	Config *config.Config
}

func NewHTTPServer(config *config.Config) *Server {
	gin.SetMode(config.GinMode)
	app := gin.Default()
	return &Server{
		Gin:    app,
		Config: config,
	}
}

func (server *Server) SetupRouter() {
	healthcheckHandler := handler.NewHealthcheckHandler()

	server.Gin.GET("/", healthcheckHandler.GetServerInformation)
	server.Gin.GET("/health", healthcheckHandler.GetHealthcheck)
}

func (server *Server) Start() {
	server.SetupRouter()
	server.Gin.Run(server.Config.ServerAddress)
}
