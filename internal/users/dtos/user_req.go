package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserReq struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserName string `json:"username"`
	Email    string  `json:"email"`
}

func (Ur *UserReq) Validate() error {
	return validation.ValidateStruct(
		Ur,
		validation.Field(&Ur.UserName, validation.Required.Error("Name is required"),
			validation.Length(5, 20).Error("Name must be 5 to 20 charectars"),
		),
		validation.Field(&Ur.Email, validation.Required.Error("Email is required"),
			is.Email.Error("Invalid email format"),
		),
		validation.Field(&Ur.Password, validation.Required.Error("Password is required"),
			validation.Length(8, 32).Error("password length must be 8 to 32 charectars"),
		),
	)
}
