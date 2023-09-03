package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBName        string `mapstructure:"DB_NAME"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string
	JWTExpiration int `mapstructure:"JWT_EXPIRATION"`
	TokenAuth     *jwtauth.JWTAuth
}

const (
	DEBUG = false
)

var config *Config

func GetConfig(path string) (*Config, error) {

	viper.SetConfigName("goham_config")
	viper.SetConfigType("env")
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
