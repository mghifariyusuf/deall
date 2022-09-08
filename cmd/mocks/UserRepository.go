// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entity "deall/cmd/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ctx, id, userID
func (_m *UserRepository) DeleteUser(ctx context.Context, id string, userID string) error {
	ret := _m.Called(ctx, id, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmailOrPhone provides a mock function with given fields: ctx, data
func (_m *UserRepository) GetUserByEmailOrPhone(ctx context.Context, data string) (*entity.User, error) {
	ret := _m.Called(ctx, data)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRoleID provides a mock function with given fields: ctx, roleID
func (_m *UserRepository) GetUserRoleID(ctx context.Context, roleID string) ([]*entity.User, error) {
	ret := _m.Called(ctx, roleID)

	var r0 []*entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.User); ok {
		r0 = rf(ctx, roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) InsertUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUser provides a mock function with given fields: ctx, page, limit
func (_m *UserRepository) ListUser(ctx context.Context, page int64, limit int64) ([]*entity.User, error) {
	ret := _m.Called(ctx, page, limit)

	var r0 []*entity.User
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []*entity.User); ok {
		r0 = rf(ctx, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, user, id
func (_m *UserRepository) UpdateUser(ctx context.Context, user *entity.User, id string) (*entity.User, error) {
	ret := _m.Called(ctx, user, id)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User, string) *entity.User); ok {
		r0 = rf(ctx, user, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.User, string) error); ok {
		r1 = rf(ctx, user, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
