// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/shipping"

	mock "github.com/stretchr/testify/mock"
)

// ShipmentItemUsecase is an autogenerated mock type for the ShipmentItemUsecase type
type ShipmentItemUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, shipment_item
func (_m *ShipmentItemUsecase) Create(c context.Context, shipment_item *domain.ShipmentItem) error {
	ret := _m.Called(c, shipment_item)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ShipmentItem) error); ok {
		r0 = rf(c, shipment_item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *ShipmentItemUsecase) CreateMany(c context.Context, items []domain.ShipmentItem) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.ShipmentItem) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *ShipmentItemUsecase) Delete(c context.Context, ID string) error {
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
func (_m *ShipmentItemUsecase) Fetch(c context.Context) ([]domain.ShipmentItem, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.ShipmentItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.ShipmentItem, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ShipmentItem); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ShipmentItem)
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
func (_m *ShipmentItemUsecase) FetchByID(c context.Context, ID string) (domain.ShipmentItem, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.ShipmentItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.ShipmentItem, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ShipmentItem); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.ShipmentItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, shipment_item
func (_m *ShipmentItemUsecase) Update(c context.Context, shipment_item *domain.ShipmentItem) error {
	ret := _m.Called(c, shipment_item)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ShipmentItem) error); ok {
		r0 = rf(c, shipment_item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewShipmentItemUsecase creates a new instance of ShipmentItemUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShipmentItemUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShipmentItemUsecase {
	mock := &ShipmentItemUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
