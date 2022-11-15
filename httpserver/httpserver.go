package httpserver

import (
	"github.com/BounkBU/doonangfangpleng/config"
	"github.com/BounkBU/doonangfangpleng/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin    *gin.Engine
	Config *config.Config
}

func NewHTTPServer(config *config.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	return &Server{
		Gin:    app,
		Config: config,
	}
}

func (server *Server) SetupRouter() {
	healthcheckHandler := handler.NewHealthcheckHandler()

	server.Gin.GET("/", healthcheckHandler.GetHealthcheck)
}

func (server *Server) Start() {
	server.SetupRouter()
	server.Gin.Run(":" + server.Config.App.Port)
}
