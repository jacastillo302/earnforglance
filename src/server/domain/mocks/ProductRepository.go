// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/catalog"

	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, product
func (_m *ProductRepository) Create(c context.Context, product *domain.Product) error {
	ret := _m.Called(c, product)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Product) error); ok {
		r0 = rf(c, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *ProductRepository) CreateMany(c context.Context, items []domain.Product) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Product) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *ProductRepository) Delete(c context.Context, ID string) error {
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
func (_m *ProductRepository) Fetch(c context.Context) ([]domain.Product, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Product, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Product); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
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
func (_m *ProductRepository) FetchByID(c context.Context, ID string) (domain.Product, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Product, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Product); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, product
func (_m *ProductRepository) Update(c context.Context, product *domain.Product) error {
	ret := _m.Called(c, product)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Product) error); ok {
		r0 = rf(c, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
