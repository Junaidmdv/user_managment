package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath("/home/junaid/user_mangment/")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Failed to load config file: ", err)
	}

	return v
}
