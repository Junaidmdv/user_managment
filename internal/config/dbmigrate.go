package config

import (
	"github.com/junaidmdv/user_mangment/internal/users/entities"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB){
   db.AutoMigrate(&entities.User{})

}