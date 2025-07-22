package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()

	err := godotenv.Load("CONFIG_PATH")

	if err != nil {
		log.Fatalf("CONFIG_PATH environment variable not set %v ", err)
	}

	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath(os.Getenv("CONFIG_PATH"))

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return v
}
