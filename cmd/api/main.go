package main

import (
	"github.com/junaidmdv/user_mangment/internal/config"
)

func main() {
	viper := config.LoadConfig()
	db := config.DbConnection(viper)
	config.Migrate(db)
	app := config.NewServer()

	config.Bootstrap(&config.BootstrapConfig{
		Viper: viper,
		DB:    db,
		App:   app,
	})
	app.Run()
}
