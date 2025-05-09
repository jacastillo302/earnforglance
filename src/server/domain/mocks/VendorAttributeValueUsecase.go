// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/vendors"

	mock "github.com/stretchr/testify/mock"
)

// VendorAttributeValueUsecase is an autogenerated mock type for the VendorAttributeValueUsecase type
type VendorAttributeValueUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, vendor_attribute_value
func (_m *VendorAttributeValueUsecase) Create(c context.Context, vendor_attribute_value *domain.VendorAttributeValue) error {
	ret := _m.Called(c, vendor_attribute_value)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.VendorAttributeValue) error); ok {
		r0 = rf(c, vendor_attribute_value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *VendorAttributeValueUsecase) CreateMany(c context.Context, items []domain.VendorAttributeValue) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.VendorAttributeValue) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *VendorAttributeValueUsecase) Delete(c context.Context, ID string) error {
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
func (_m *VendorAttributeValueUsecase) Fetch(c context.Context) ([]domain.VendorAttributeValue, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.VendorAttributeValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.VendorAttributeValue, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.VendorAttributeValue); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.VendorAttributeValue)
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
func (_m *VendorAttributeValueUsecase) FetchByID(c context.Context, ID string) (domain.VendorAttributeValue, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.VendorAttributeValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.VendorAttributeValue, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.VendorAttributeValue); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.VendorAttributeValue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, vendor_attribute_value
func (_m *VendorAttributeValueUsecase) Update(c context.Context, vendor_attribute_value *domain.VendorAttributeValue) error {
	ret := _m.Called(c, vendor_attribute_value)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.VendorAttributeValue) error); ok {
		r0 = rf(c, vendor_attribute_value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewVendorAttributeValueUsecase creates a new instance of VendorAttributeValueUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVendorAttributeValueUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *VendorAttributeValueUsecase {
	mock := &VendorAttributeValueUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
