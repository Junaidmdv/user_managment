package usecase

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/junaidmdv/user_mangment/internal/users"
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
)

type UserUsecase struct {
	Validate   validation.Validatable
	Repository users.Respository
}

func NewUser(rp users.Respository, v validation.Validatable) *UserUsecase {
	return &UserUsecase{
		Repository: rp,
		Validate:   v,
	}
}

func (uc *UserUsecase) Signup(user *dtos.UserReq) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := uc.Repository.IsEmailExist(user.Email); err != nil {
		return err
	}

	if err := uc.Repository.AddUser(user); err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) GetUsers() ([]dtos.UserResponse, error) {
	users, err := uc.Repository.GetUsers()
	if err != nil {
		return []dtos.UserResponse{}, nil
	}

	return users, nil
}
