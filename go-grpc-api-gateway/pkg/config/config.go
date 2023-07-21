package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	err = godotenv.Load("./pkg/config/envs/dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&c)

	return
}
