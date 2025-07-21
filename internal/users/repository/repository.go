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

func (ur *UserRepository) IsEmailExist(email string) bool {
	var exists bool
	ur.db.Model(&ur.user).
		Select("count(*)>0").
		Where("email=?", email).Find(&exists)
	return exists
}

func (ur *UserRepository) GetUsers() (*[]dtos.UserResponse, error) {
	var user []dtos.UserResponse
	if err := ur.db.Model(entities.User{}).Find(&user).Error; err != nil {
		return nil, entities.ErrDbFailure
	}
	return &user, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	result := ur.db.Delete(&entities.User{}, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entities.ErrUserNotfound
		} else {
			return entities.ErrInternalServerError
		}
	}
	return nil
}

func (ur *UserRepository) UpdateUser(id int, user *dtos.UserResponse) error {
	updateuser := make(map[string]interface{})

	if user.UserName != "" {
		updateuser["user_name"] = user.UserName
	}

	if user.Email != "" {
		updateuser["user_name"] = user.UserName
	}

	result := ur.db.Model(entities.User{}).Where("id=?", id).Updates(updateuser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entities.ErrUserNotfound
		} else {
			return entities.ErrInternalServerError
		}
	}
	return nil
}
