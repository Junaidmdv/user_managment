package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
	"github.com/junaidmdv/user_mangment/internal/users/entities"
	"github.com/junaidmdv/user_mangment/internal/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	var user dtos.UserReq
	err := faker.FakeData(&user)
	assert.NoError(t, err)
	mockuc := mocks.NewMockUsecase(t)
	// mockuc.On("Signup", mock.Anything).Return(nil).Once()

	t.Run("error_responce", func(t *testing.T) {
		mockuc.On("Signup", mock.Anything).Return(entities.ErrDbFailure).Once()
		handler := NewHandler(mockuc)
		expected := ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: entities.ErrDbFailure.Error(),
		}
		r := gin.New()
		r.POST("/signup", handler.Signup)
		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responseData, _ := io.ReadAll(w.Body)

		var actualError ErrorResponse
		err := json.Unmarshal(responseData, &actualError)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expected, actualError)

	})

	t.Run("success", func(t *testing.T) {
		mockuc.On("Signup", mock.Anything).Return(nil).Once()
		handler := NewHandler(mockuc)
		r := gin.New()
		r.POST("signup", handler.Signup)
		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetUsers(t *testing.T) {
	var users *[]dtos.UserResponse
	err := faker.FakeData(&users)
	assert.NoError(t, err)
	mockuc := mocks.NewMockUsecase(t)

	t.Run("error_case", func(t *testing.T) {
		expected := ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: entities.ErrDbFailure.Error(),
		}
		mockuc.On("GetUsers").Return(&[]dtos.UserResponse{}, entities.ErrDbFailure).Once()
		handler := NewHandler(mockuc)
		r := gin.New()
		r.GET("/users", handler.GetUsers)
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responceData, _ := io.ReadAll(w.Body)
		var actual ErrorResponse
		err := json.Unmarshal(responceData, &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)

	})

	t.Run("success", func(t *testing.T) {
		userData := make(map[string]interface{})
		mockuc.On("GetUsers").Return(users, nil).Once()
		handler := NewHandler(mockuc)
		r := gin.New()
		r.GET("/users", handler.GetUsers)
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responceData, _ := io.ReadAll(w.Body)
		err := json.Unmarshal(responceData, &userData)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, userData["users"])
	})
}

func TestDeleteUser(t *testing.T) {
	mockuc := mocks.NewMockUsecase(t)

	t.Run("error", func(t *testing.T) {
		expected := ErrorResponse{
			Code:    http.StatusNotFound,
			Message: entities.ErrUserNotfound.Error(),
		}
		mockuc.On("DeleteUser", mock.AnythingOfType("int")).Return(entities.ErrUserNotfound).Once()
		handler := NewHandler(mockuc)
		r := gin.New()
		ID := `234`
		r.DELETE("/users/:id", handler.DeleteUser)
		req, _ := http.NewRequest(http.MethodDelete, "/users/"+ID, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responceData, _ := io.ReadAll(w.Body)
		var actual ErrorResponse
		err := json.Unmarshal(responceData, &actual)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expected, actual)

	})

	t.Run("success", func(t *testing.T) {
		mockuc.On("DeleteUser", mock.AnythingOfType("int")).Return(nil).Once()
		handler := NewHandler(mockuc)
		r := gin.New()
		ID := `234`
		r.DELETE("/users/:id", handler.DeleteUser)
		req, _ := http.NewRequest(http.MethodDelete, "/users/"+ID, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestUpdateUser(t *testing.T) {
	mockuc := mocks.NewMockUsecase(t)

	t.Run("error", func(t *testing.T) {
		expected := ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: entities.ErrUserNotfound.Error(),
		}
		handler := NewHandler(mockuc)
		mockuc.On("UpdateUser", mock.Anything, mock.Anything).Return(entities.ErrUserNotfound).Once()
		r := gin.New()
		ID := `234`
		r.PATCH("/users/:id", handler.UpdateUser)
		req, _ := http.NewRequest(http.MethodPatch, "/users/"+ID, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responceData, _ := io.ReadAll(w.Body)
		var actual ErrorResponse
		err := json.Unmarshal(responceData, &actual)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expected, actual)

	})

	t.Run("success", func(t *testing.T) {
		mockuc.On("UpdateUser", mock.AnythingOfType("int"), mock.Anything).Return(nil).Once()
		handler := NewHandler(mockuc)
		r := gin.New()
		ID := `2`
		r.PATCH("/users/:id", handler.UpdateUser)
		req, _ := http.NewRequest(http.MethodPatch, "/users/"+ID, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

	})
}
