// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/stores"

	mock "github.com/stretchr/testify/mock"
)

// StoreUsecase is an autogenerated mock type for the StoreUsecase type
type StoreUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, store
func (_m *StoreUsecase) Create(c context.Context, store *domain.Store) error {
	ret := _m.Called(c, store)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Store) error); ok {
		r0 = rf(c, store)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *StoreUsecase) CreateMany(c context.Context, items []domain.Store) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Store) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *StoreUsecase) Delete(c context.Context, ID string) error {
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
func (_m *StoreUsecase) Fetch(c context.Context) ([]domain.Store, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.Store
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Store, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Store); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Store)
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
func (_m *StoreUsecase) FetchByID(c context.Context, ID string) (domain.Store, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.Store
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Store, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Store); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.Store)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, store
func (_m *StoreUsecase) Update(c context.Context, store *domain.Store) error {
	ret := _m.Called(c, store)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Store) error); ok {
		r0 = rf(c, store)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStoreUsecase creates a new instance of StoreUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStoreUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *StoreUsecase {
	mock := &StoreUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
