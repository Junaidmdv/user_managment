package users

import (
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
)

type Respository interface {
	AddUser(user *dtos.UserReq) error
	IsEmailExist(string) bool
	GetUsers() (*[]dtos.UserResponse, error)
	DeleteUser(int)error
	UpdateUser(int,*dtos.UserResponse)error
}
