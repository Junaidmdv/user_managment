package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/junaidmdv/user_mangment/internal/users"
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
	"github.com/junaidmdv/user_mangment/internal/users/entities"
)

type UserHandler struct {
	UseCase users.UsercaseLayer
}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHandler(uc users.UsercaseLayer) *UserHandler {
	return &UserHandler{
		UseCase: uc,
	}
}

func (U *UserHandler) Signup(c *gin.Context) {
	var user dtos.UserReq

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest,
			ErrorResponse{Code: http.StatusBadRequest,
				Message: entities.ErrInvalidRequestBody.Error()})
		return
	}
	if err := U.UseCase.Signup(&user); err != nil {
		c.JSON(getstatusCode(err), ErrorResponse{Code: getstatusCode(err), Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user singup successfull",
	})

}

func (U *UserHandler) GetUsers(c *gin.Context) {
	users, err := U.UseCase.GetUsers()
	if err != nil {
		c.JSON(getstatusCode(err), ErrorResponse{
			Code:    getstatusCode(err),
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user data retrieved",
		"users":   users,
	})

}

func (U *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(getstatusCode(entities.ErrBadParamInput),
			ErrorResponse{Code: getstatusCode(entities.ErrBadParamInput),
				Message: entities.ErrBadParamInput.Error(),
			},
		)
		return
	}
	if err = U.UseCase.DeleteUser(id); err != nil {
		c.JSON(getstatusCode(err), ErrorResponse{
			Code:    getstatusCode(err),
			Message: err.Error(),
		},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user account deleted",
	})

}

func (U *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(getstatusCode(entities.ErrBadParamInput),
			ErrorResponse{Code: getstatusCode(entities.ErrBadParamInput),
				Message: entities.ErrBadParamInput.Error(),
			},
		)
		return
	}
	var user dtos.UserResponse

	if err = U.UseCase.UpdateUser(id, &user); err != nil {
		c.JSON(getstatusCode(entities.ErrBadParamInput),
			ErrorResponse{Code: getstatusCode(entities.ErrBadParamInput),
				Message: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user account updated",
	})
}

func getstatusCode(err error) int {

	switch err {
	case entities.ErrBadParamInput:
		return http.StatusBadRequest
	case entities.ErrBadReqBody:
		return http.StatusBadRequest
	case entities.ErrConflict:
		return http.StatusConflict
	case entities.ErrInternalServerError:
		return http.StatusInternalServerError
	case entities.ErrNotFound:
		return http.StatusNotFound
	case entities.ErrEmailExist:
		return http.StatusBadRequest
	case entities.ErrDbFailure:
		return http.StatusInternalServerError
	case entities.ErrUserNotfound:
		return http.StatusNotFound
	}

	return http.StatusUnprocessableEntity
}
