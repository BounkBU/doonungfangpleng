package main

import (
	"fmt"

	"github.com/BounkBU/doonungfangpleng/config"
	"github.com/BounkBU/doonungfangpleng/httpserver"
	"github.com/BounkBU/doonungfangpleng/pkg/database"
	"github.com/BounkBU/doonungfangpleng/pkg/logger"
	"github.com/BounkBU/doonungfangpleng/pkg/util"
	log "github.com/sirupsen/logrus"
)

var serverConfig *config.Config

func init() {
	serverConfig = config.LoadConfig()

	logger.InitLogger(serverConfig.App)
}

func main() {
	token, err := util.SpotifyConnection("5bac8f81d64f43058a6d033152ce770b", "fe77ea1ab76f42fc81f5f3a0b2b841dc")
	if err != nil {
		log.Fatalf("error, create spotify token, %s", err.Error())
	}
	fmt.Println(token.AccessToken)

	db, err := database.NewMySQLDatabaseConnection(serverConfig)
	if err != nil {
		log.Fatalf("error, create mysql database connection, %s", err.Error())
	}

	server := httpserver.NewHTTPServer(serverConfig, db)

	server.Start()
}
