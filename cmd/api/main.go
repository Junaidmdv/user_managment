package main

import (
	"log"

	"github.com/junaidmdv/user_mangment/internal/config"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath("../../")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Failed to load config file: ", err)
	}

	db := config.DbConnection(v)
	config.Migrate(db)
	app := config.NewServer()

	config.Bootstrap(&config.BootstrapConfig{
		Viper: v,
		DB:    db,
		App:   app,
	})
	app.Run(":8000")
}
