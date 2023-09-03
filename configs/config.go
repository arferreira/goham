package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DBName        string
	DBUser        string
	DBPass        string
	DBHost        string
	DBPort        string
	WebServerPort string
	JWTSecret     string
	JWTExpiration int
	TokenAuth     *jwtauth.JWTAuth
}

const (
	DEBUG = true
)

var config *Config

func GetConfig(path string) (*Config, error) {
	// TODO: Read config file and return config struct
	viper.SetConfigName("goham_config")
	viper.SetConfigType(".env")
	viper.AddConfigPath(path)

	if DEBUG {
		viper.SetConfigName(".env")
	} else {
		viper.SetConfigName(".env.production")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)

	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	return config, err
}
