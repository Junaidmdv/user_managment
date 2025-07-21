package usecase

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
	"github.com/junaidmdv/user_mangment/internal/users/entities"
	"github.com/junaidmdv/user_mangment/internal/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	mockrepo := mocks.NewMockRepository(t)
	t.Run("validation_error", func(t *testing.T) {
		user := &dtos.UserReq{
			UserName: "junaid",
			Email:    "junaidgmail.com",
			Password: "1234567",
		}

		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		got := uc.Signup(user)
		assert.Error(t, got)
		mockrepo.AssertExpectations(t)
	})

	t.Run("error_email_exist", func(t *testing.T) {
		user := &dtos.UserReq{
			UserName: "junaid",
			Email:    "junaid@gmail.com",
			Password: "12345678",
		}
		mockrepo.On("IsEmailExist", user.Email).Return(true).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		got := uc.Signup(user)
		assert.Error(t, got)
		mockrepo.AssertExpectations(t)

	})

	t.Run("success", func(t *testing.T) {
		user := &dtos.UserReq{
			UserName: "junaid",
			Email:    "junaid@gmail.com",
			Password: "12345678",
		}

		mockrepo.On("IsEmailExist", user.Email).Return(false).Once()
		mockrepo.On("AddUser", mock.Anything).Return(nil).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		got := uc.Signup(user)

		assert.NoError(t, got)

		mockrepo.AssertExpectations(t)

	})
}

func TestGetUsers(t *testing.T) {
	mockrepo := mocks.NewMockRepository(t)

	users := make([]dtos.UserResponse, 0)

	

	t.Run("datbase_error", func(t *testing.T) {
		mockrepo.On("GetUsers").Return(nil, entities.ErrDbFailure).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		usr, err := uc.GetUsers()
		assert.Equal(t, &users,usr)
		assert.Error(t, err)
		mockrepo.AssertExpectations(t)

	})

	t.Run("success", func(t *testing.T) {
		mockrepo.On("GetUsers").Return(&users, nil).Once()
		user := dtos.UserResponse{
		UserName: "junaid",
		Email:    "junaid@gmail.com",
	    }
	     users = append(users, user)
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		usr, err := uc.GetUsers()
		assert.NotNil(t, usr)
		assert.Nil(t, err)

	})

}

func TestDeleteUser(t *testing.T) {
    mockrepo := mocks.NewMockRepository(t)

	t.Run("error_case",func(t *testing.T) {
        mockrepo.On("DeleteUser",mock.AnythingOfType("int")).Return(entities.ErrNotFound).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		id:=int(1)
		err:= uc.DeleteUser(id)

		assert.Equal(t,entities.ErrNotFound,err)
		
        
	})
	t.Run("success",func(t *testing.T) {
        mockrepo.On("DeleteUser",mock.AnythingOfType("int")).Return(nil).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		id:=int(1)
		err:= uc.DeleteUser(id)
        assert.Nil(t,err)
        mockrepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	 mockrepo := mocks.NewMockRepository(t)

	 t.Run("error_case",func(t *testing.T) {
		mockrepo.On("UpdateUser",mock.Anything,mock.Anything).Return(entities.ErrNotFound).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		id:=int(1)
		user := &dtos.UserResponse{
          UserName: "junaid",
		  Email: "junaid@gmail.com",
        }
		err:= uc.UpdateUser(id,user)
        assert.Equal(t,entities.ErrNotFound,err)
        mockrepo.AssertExpectations(t)
	 })


	 t.Run("success",func(t *testing.T) {
		mockrepo.On("UpdateUser",mock.Anything,mock.Anything).Return(nil).Once()
		var mockvalidate validation.Validatable
		uc := NewUser(mockrepo, mockvalidate)
		id:=int(1)
		
		user := &dtos.UserResponse{
          UserName: "junaid",
		  Email: "junaid@gmail.com",
        }
		err:= uc.UpdateUser(id,user)
        assert.Nil(t,err)
        mockrepo.AssertExpectations(t)
	 })

}
