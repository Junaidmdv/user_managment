package repository

import (
	"fmt"

	"github.com/junaidmdv/user_mangment/internal/users/dtos"
	"github.com/junaidmdv/user_mangment/internal/users/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db   *gorm.DB
	user entities.User
}

func NewUserRepository(db *gorm.DB, user entities.User) *UserRepository {
	return &UserRepository{
		db:   db,
		user: user,
	}
}

func (ur *UserRepository) AddUser(user *dtos.UserReq) error {
	nuser := entities.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := ur.db.Create(&nuser).Error; err != nil {
		fmt.Println(err)
		return entities.ErrDbFailure
	}
	return nil
}

func (ur *UserRepository) IsEmailExist(email string) error {
	var exists bool
	ur.db.Model(&ur.user).
		Select("count(*)>0").
		Where("email=?", email).Find(&exists)
	if exists {
		return entities.ErrEmailExist
	}
	return nil
}

func (ur *UserRepository) GetUsers() ([]dtos.UserResponse, error) {
	var user []dtos.UserResponse
	if err := ur.db.Model(entities.User{}).Find(&user).Error; err != nil {
		return nil, entities.ErrDbFailure
	}
	return user, nil
}
