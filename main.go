package main

import (
	"log"

	"github.com/BounkBU/doonangfangpleng/config"
	"github.com/BounkBU/doonangfangpleng/httpserver"
)

var appConfig *config.Config
var err error

func init() {
	appConfig, err = config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := httpserver.NewHTTPServer(appConfig)

	server.Start()
}
