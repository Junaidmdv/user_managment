package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbconfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	DB_name  string `mapstructure:"DB_NAME"`
	DB_port  string `mapstructure:"DB_PORT"`
}

func DbConnection(config *viper.Viper) *gorm.DB {

	var db Dbconfig
	if err := config.Unmarshal(&db); err != nil {
		log.Fatal("Failed to unmarshal env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", db.Host, db.Username, db.Password, db.DB_name, db.DB_port)

	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("databse configration failure", err)
	}
	return Db

}
