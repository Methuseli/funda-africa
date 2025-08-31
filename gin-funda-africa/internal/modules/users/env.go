package users

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`

	// GoogleClientID         string `mapstructure:"GOOGLE_CLIENT_ID"`
	// GoogleClientSecret     string `mapstructure:"GOOGLE_SECRET_KEY"`
	// GoogleOAuthRedirectUrl string `mapstructure:"GOOGLE_REDIRECT_URL"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	env := os.Getenv("GIN_MODE")

	if env == "" || env == "development" {
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName("dev")

		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
		return
	} else {
		return
	}


}