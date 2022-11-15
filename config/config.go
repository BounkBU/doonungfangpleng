package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App App
}

type App struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	var appConfig App
	appConfig.Port = os.Getenv("PORT")

	return &Config{
		App: appConfig,
	}
}
