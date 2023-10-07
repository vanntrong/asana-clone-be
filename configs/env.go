package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	DBUrl              string `mapstructure:"DB_URL"`
	PORT               string `mapstructure:"PORT"`
	GoogleClientId     string `mapstructure:"GOOGLE_CLIENT_ID"`
}

var AppConfig Config

func LoadEnv(path string, config *Config) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
