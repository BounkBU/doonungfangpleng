package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Database Database
}

type App struct {
	Env           string `mapstructure:"ENV"`
	GinMode       string `mapstructure:"GIN_MODE" default:"release"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type Database struct {
	Hostname string `mapstructure:"MYSQL_HOSTNAME"`
	Port     string `mapstructure:"MYSQL_PORT"`
	Username string `mapstructure:"MYSQL_USERNAME"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
	Database string `mapstructure:"MYSQL_DATABASE"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	var appConfig App
	var databaseConfig Database
	var config *Config

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error, can't read config in .env file, %s", err.Error())
		return config, err
	}

	if err = viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("error, can't parse app config, %s", err.Error())
		return config, err
	}

	if err = viper.Unmarshal(&databaseConfig); err != nil {
		log.Fatalf("error, can't parse database config, %s", err.Error())
		return config, err
	}

	config = &Config{
		App:      appConfig,
		Database: databaseConfig,
	}

	log.Println("Load config from .env file successfully")
	return config, nil
}
