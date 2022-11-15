package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
}

type App struct {
	Env     string `mapstructure:"ENV"`
	GinMode string `mapstructure:"GIN_MODE" default:"release"`
	Port    string `mapstructure:"PORT"`
}

type Database struct {
	Hostname string `mapstructure:"MYSQL_HOSTNAME"`
	Port     string `mapstructure:"MYSQL_PORT"`
	Username string `mapstructure:"MYSQL_USERNAME"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
	Database string `mapstructure:"MYSQL_DATABASE"`
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	var appConfig App
	var databaseConfig Database

	appConfig.Env = os.Getenv("ENV")
	appConfig.GinMode = os.Getenv("GIN_MODE")
	appConfig.Port = os.Getenv("PORT")

	databaseConfig.Hostname = os.Getenv("MYSQL_HOSTNAME")
	databaseConfig.Port = os.Getenv("MYSQL_PORT")
	databaseConfig.Username = os.Getenv("MYSQL_USERNAME")
	databaseConfig.Password = os.Getenv("MYSQL_PASSWORD")
	databaseConfig.Database = os.Getenv("MYSQL_DATABASE")

	return &Config{
		App:      appConfig,
		Database: databaseConfig,
	}
}

// func LoadConfig(path string) (*Config, error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigFile(".env")

// 	viper.AutomaticEnv()

// 	var appConfig App
// 	var databaseConfig Database
// 	var config *Config

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Errorf("error, can't read config in .env file, %s", err.Error())
// 		return config, err
// 	}

// 	if err = viper.Unmarshal(&appConfig); err != nil {
// 		log.Errorf("error, can't parse app config, %s", err.Error())
// 		return config, err
// 	}

// 	if err = viper.Unmarshal(&databaseConfig); err != nil {
// 		log.Errorf("error, can't parse database config, %s", err.Error())
// 		return config, err
// 	}

// 	config = &Config{
// 		App:      appConfig,
// 		Database: databaseConfig,
// 	}

// 	log.Println("Load config from .env file successfully")
// 	return config, nil
// }
