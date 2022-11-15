package config

import (
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
		return config, err
	}

	if err = viper.Unmarshal(&appConfig); err != nil {
		return config, err
	}

	if err = viper.Unmarshal(&databaseConfig); err != nil {
		return config, err
	}

	config = &Config{
		App:      appConfig,
		Database: databaseConfig,
	}

	return config, nil
}
