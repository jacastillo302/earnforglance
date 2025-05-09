// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/forums"

	mock "github.com/stretchr/testify/mock"
)

// ForumSettingsRepository is an autogenerated mock type for the ForumSettingsRepository type
type ForumSettingsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, forum_settings
func (_m *ForumSettingsRepository) Create(c context.Context, forum_settings *domain.ForumSettings) error {
	ret := _m.Called(c, forum_settings)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ForumSettings) error); ok {
		r0 = rf(c, forum_settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *ForumSettingsRepository) CreateMany(c context.Context, items []domain.ForumSettings) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.ForumSettings) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *ForumSettingsRepository) Delete(c context.Context, ID string) error {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: c
func (_m *ForumSettingsRepository) Fetch(c context.Context) ([]domain.ForumSettings, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.ForumSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.ForumSettings, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ForumSettings); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ForumSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: c, ID
func (_m *ForumSettingsRepository) FetchByID(c context.Context, ID string) (domain.ForumSettings, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.ForumSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.ForumSettings, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ForumSettings); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.ForumSettings)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, forum_settings
func (_m *ForumSettingsRepository) Update(c context.Context, forum_settings *domain.ForumSettings) error {
	ret := _m.Called(c, forum_settings)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ForumSettings) error); ok {
		r0 = rf(c, forum_settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewForumSettingsRepository creates a new instance of ForumSettingsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewForumSettingsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ForumSettingsRepository {
	mock := &ForumSettingsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
