// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "docker/domain"
	request "docker/domain/web/request"

	mock "github.com/stretchr/testify/mock"

	response "docker/domain/web/response"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *UserUsecase) Create(_a0 request.UserCreateRequest) (*domain.User, error) {
	ret := _m.Called(_a0)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(request.UserCreateRequest) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(request.UserCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: _a0
func (_m *UserUsecase) Login(_a0 request.LoginRequest) (*response.SuccessLogin, error) {
	ret := _m.Called(_a0)

	var r0 *response.SuccessLogin
	if rf, ok := ret.Get(0).(func(request.LoginRequest) *response.SuccessLogin); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.SuccessLogin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(request.LoginRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields:
func (_m *UserUsecase) ReadAll() (*domain.Users, error) {
	ret := _m.Called()

	var r0 *domain.Users
	if rf, ok := ret.Get(0).(func() *domain.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadByID provides a mock function with given fields: id
func (_m *UserUsecase) ReadByID(id int) (*domain.User, error) {
	ret := _m.Called(id)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(int) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
