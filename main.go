package main

import (
	"github.com/BounkBU/doonangfangpleng/config"
	"github.com/BounkBU/doonangfangpleng/httpserver"
)

var appConfig *config.Config

func init() {
	appConfig = config.LoadConfig()
}

func main() {
	server := httpserver.NewHTTPServer(appConfig)

	server.Start()
}
