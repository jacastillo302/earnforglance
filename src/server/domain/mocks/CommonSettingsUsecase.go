// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/common"

	mock "github.com/stretchr/testify/mock"
)

// CommonSettingsUsecase is an autogenerated mock type for the CommonSettingsUsecase type
type CommonSettingsUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, common_settings
func (_m *CommonSettingsUsecase) Create(c context.Context, common_settings *domain.CommonSettings) error {
	ret := _m.Called(c, common_settings)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CommonSettings) error); ok {
		r0 = rf(c, common_settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *CommonSettingsUsecase) CreateMany(c context.Context, items []domain.CommonSettings) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.CommonSettings) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *CommonSettingsUsecase) Delete(c context.Context, ID string) error {
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
func (_m *CommonSettingsUsecase) Fetch(c context.Context) ([]domain.CommonSettings, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.CommonSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.CommonSettings, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.CommonSettings); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CommonSettings)
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
func (_m *CommonSettingsUsecase) FetchByID(c context.Context, ID string) (domain.CommonSettings, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.CommonSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.CommonSettings, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.CommonSettings); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.CommonSettings)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, common_settings
func (_m *CommonSettingsUsecase) Update(c context.Context, common_settings *domain.CommonSettings) error {
	ret := _m.Called(c, common_settings)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CommonSettings) error); ok {
		r0 = rf(c, common_settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommonSettingsUsecase creates a new instance of CommonSettingsUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommonSettingsUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommonSettingsUsecase {
	mock := &CommonSettingsUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
