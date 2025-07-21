package mocks

import (
	"github.com/junaidmdv/user_mangment/internal/users/dtos"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (_m *MockRepository) AddUser(user *dtos.UserReq) error {
	ret := _m.Called(user)

	var r0 error

	if rf, ok := ret.Get(0).(func(*dtos.UserReq) error); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Error(0)
		}
	}

	return r0

}

func (_m *MockRepository) IsEmailExist(email string) bool {
	ret := _m.Called(email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Bool(0)
		}
	}

	return r0
}

func (_m *MockRepository) GetUsers() (*[]dtos.UserResponse, error) {
	ret := _m.Called()

	var r0 *[]dtos.UserResponse

	if rf, ok := ret.Get(0).(func() *[]dtos.UserResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]dtos.UserResponse)
		}
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		if ret.Error(1) != nil {
			r1 = ret.Error(1)
		}
	}

	return r0, r1

}

func (_m *MockRepository) DeleteUser(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		if ret.Error(0) != nil {
			r0 = ret.Error(0)
		}
	}

	return r0
}

func (_m *MockRepository) UpdateUser(id int, user *dtos.UserResponse) error {
	ret := _m.Called(id, user)

	var r0 error

	if rf, ok := ret.Get(0).(func(int, *dtos.UserResponse) error); ok {
		r0 = rf(id, user)
	} else {
		if ret.Error(0) != nil {
			r0 = ret.Error(0)
		}

	}
	return r0
}

func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
