package config

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	http "github.com/junaidmdv/user_mangment/internal/users/delivery/http"
	"github.com/junaidmdv/user_mangment/internal/users/entities"
	"github.com/junaidmdv/user_mangment/internal/users/repository"
	"github.com/junaidmdv/user_mangment/internal/users/usecase"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB    *gorm.DB
	App   *gin.Engine
	Viper *viper.Viper
}

func Bootstrap(b *BootstrapConfig) {

	//validation
	var validation validation.Validatable

	//users
	User_repository := repository.NewUserRepository(b.DB, entities.User{})
	User_usecase := usecase.NewUser(User_repository, validation)
	User_handler := http.NewHandler(User_usecase)

	routes := http.RouteConfig{
		App:            b.App,
		UserController: User_handler,
	}
	routes.Setup()

}
