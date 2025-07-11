package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB    *gorm.DB
	App   *gin.Engine
	Viper *viper.Viper
}


func Bootstrap( b *BootstrapConfig){

}


