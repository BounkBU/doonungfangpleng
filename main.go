package main

import (
	"github.com/BounkBU/doonungfangpleng/config"
	"github.com/BounkBU/doonungfangpleng/httpserver"
	"github.com/BounkBU/doonungfangpleng/pkg/database"
	"github.com/BounkBU/doonungfangpleng/pkg/logger"
	log "github.com/sirupsen/logrus"
)

var serverConfig *config.Config

func init() {
	serverConfig = config.LoadConfig()

	logger.InitLogger(serverConfig.App)
}

func main() {
	db, err := database.NewMySQLDatabaseConnection(serverConfig)
	if err != nil {
		log.Fatalf("error, create mysql database connection, %s", err.Error())
	}

	server := httpserver.NewHTTPServer(serverConfig, db)

	server.Start()
}
