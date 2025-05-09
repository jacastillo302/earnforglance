// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/shipping"

	mock "github.com/stretchr/testify/mock"
)

// PickupPointRepository is an autogenerated mock type for the PickupPointRepository type
type PickupPointRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, pickup_point
func (_m *PickupPointRepository) Create(c context.Context, pickup_point *domain.PickupPoint) error {
	ret := _m.Called(c, pickup_point)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.PickupPoint) error); ok {
		r0 = rf(c, pickup_point)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *PickupPointRepository) CreateMany(c context.Context, items []domain.PickupPoint) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.PickupPoint) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *PickupPointRepository) Delete(c context.Context, ID string) error {
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
func (_m *PickupPointRepository) Fetch(c context.Context) ([]domain.PickupPoint, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.PickupPoint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.PickupPoint, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.PickupPoint); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PickupPoint)
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
func (_m *PickupPointRepository) FetchByID(c context.Context, ID string) (domain.PickupPoint, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.PickupPoint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.PickupPoint, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.PickupPoint); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.PickupPoint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, pickup_point
func (_m *PickupPointRepository) Update(c context.Context, pickup_point *domain.PickupPoint) error {
	ret := _m.Called(c, pickup_point)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.PickupPoint) error); ok {
		r0 = rf(c, pickup_point)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPickupPointRepository creates a new instance of PickupPointRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPickupPointRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PickupPointRepository {
	mock := &PickupPointRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
