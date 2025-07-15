package users

import (
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
)

type Respository interface {
	AddUser(user *dtos.UserReq) error
	IsEmailExist(string) error
	GetUsers() ([]dtos.UserResponse, error)
}
