package entities

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID          uint `gorm:"primarykey"`
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
