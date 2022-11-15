package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env            string `mapstructure:"ENV"`
	GinMode        string `mapstructure:"GIN_MODE" default:"release"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	MYSQL_HOSTNAME string `mapstructure:"MYSQL_HOSTNAME"`
	MYSQL_PORT     string `mapstructure:"MYSQL_PORT"`
	MYSQL_USERNAME string `mapstructure:"MYSQL_USERNAME"`
	MYSQL_PASSWORD string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_DATABASE string `mapstructure:"MYSQL_DATABASE"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
