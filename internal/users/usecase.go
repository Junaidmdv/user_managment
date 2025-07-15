package users

import (
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
)

type UsercaseLayer interface {
	Signup(*dtos.UserReq) error
	GetUsers() ([]dtos.UserResponse, error)
}
