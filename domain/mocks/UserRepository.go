// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "docker/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CheckLogin provides a mock function with given fields: user
func (_m *UserRepository) CheckLogin(user *domain.User) (*domain.User, bool, error) {
	ret := _m.Called(user)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*domain.User) bool); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*domain.User) error); ok {
		r2 = rf(user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}


// Create provides a mock function with given fields: user
func (_m *UserRepository) Create(user *domain.User) (*domain.User, error) {
	ret := _m.Called(user)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields:
func (_m *UserRepository) ReadAll() (*domain.Users, error) {
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
func (_m *UserRepository) ReadByID(id int) (*domain.User, error) {
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
