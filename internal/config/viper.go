package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig()*viper.Viper {
	  v:= viper.New()
	  v.SetConfigName("config")
	  v.SetConfigType("env")
	  v.AddConfigPath("./../../")
	  v.AddConfigPath("/config")
	  if err := v.ReadInConfig();err != nil{
		   log.Fatal(err)
	  }

	 return v

}
