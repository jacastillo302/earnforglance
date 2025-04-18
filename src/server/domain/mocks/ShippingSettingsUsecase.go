// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/shipping"

	mock "github.com/stretchr/testify/mock"
)

// ShippingSettingsUsecase is an autogenerated mock type for the ShippingSettingsUsecase type
type ShippingSettingsUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, shipping_settings
func (_m *ShippingSettingsUsecase) Create(c context.Context, shipping_settings *domain.ShippingSettings) error {
	ret := _m.Called(c, shipping_settings)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ShippingSettings) error); ok {
		r0 = rf(c, shipping_settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *ShippingSettingsUsecase) CreateMany(c context.Context, items []domain.ShippingSettings) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.ShippingSettings) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *ShippingSettingsUsecase) Delete(c context.Context, ID string) error {
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
func (_m *ShippingSettingsUsecase) Fetch(c context.Context) ([]domain.ShippingSettings, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.ShippingSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.ShippingSettings, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ShippingSettings); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ShippingSettings)
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
func (_m *ShippingSettingsUsecase) FetchByID(c context.Context, ID string) (domain.ShippingSettings, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.ShippingSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.ShippingSettings, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ShippingSettings); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.ShippingSettings)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, shipping_settings
func (_m *ShippingSettingsUsecase) Update(c context.Context, shipping_settings *domain.ShippingSettings) error {
	ret := _m.Called(c, shipping_settings)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ShippingSettings) error); ok {
		r0 = rf(c, shipping_settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewShippingSettingsUsecase creates a new instance of ShippingSettingsUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShippingSettingsUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShippingSettingsUsecase {
	mock := &ShippingSettingsUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
