package main

import (
	"log"

	"github.com/BounkBU/doonungfangpleng/config"
	"github.com/BounkBU/doonungfangpleng/httpserver"
	"github.com/BounkBU/doonungfangpleng/pkg/database"
	"github.com/BounkBU/doonungfangpleng/pkg/logger"
)

var serverConfig *config.Config
var err error

func init() {
	serverConfig, err = config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	logger.InitLogger(serverConfig.App)
}

func main() {
	db, err := database.NewMySQLDatabaseConnection(serverConfig.Database)
	if err != nil {
		log.Fatal(err)
	}

	server := httpserver.NewHTTPServer(serverConfig, db)

	server.Start()
}
