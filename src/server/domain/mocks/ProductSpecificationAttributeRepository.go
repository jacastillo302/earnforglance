// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/catalog"

	mock "github.com/stretchr/testify/mock"
)

// ProductSpecificationAttributeRepository is an autogenerated mock type for the ProductSpecificationAttributeRepository type
type ProductSpecificationAttributeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, product_specification_attribute
func (_m *ProductSpecificationAttributeRepository) Create(c context.Context, product_specification_attribute *domain.ProductSpecificationAttribute) error {
	ret := _m.Called(c, product_specification_attribute)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ProductSpecificationAttribute) error); ok {
		r0 = rf(c, product_specification_attribute)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *ProductSpecificationAttributeRepository) CreateMany(c context.Context, items []domain.ProductSpecificationAttribute) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.ProductSpecificationAttribute) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *ProductSpecificationAttributeRepository) Delete(c context.Context, ID string) error {
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
func (_m *ProductSpecificationAttributeRepository) Fetch(c context.Context) ([]domain.ProductSpecificationAttribute, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.ProductSpecificationAttribute
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.ProductSpecificationAttribute, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ProductSpecificationAttribute); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProductSpecificationAttribute)
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
func (_m *ProductSpecificationAttributeRepository) FetchByID(c context.Context, ID string) (domain.ProductSpecificationAttribute, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.ProductSpecificationAttribute
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.ProductSpecificationAttribute, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ProductSpecificationAttribute); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.ProductSpecificationAttribute)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, product_specification_attribute
func (_m *ProductSpecificationAttributeRepository) Update(c context.Context, product_specification_attribute *domain.ProductSpecificationAttribute) error {
	ret := _m.Called(c, product_specification_attribute)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ProductSpecificationAttribute) error); ok {
		r0 = rf(c, product_specification_attribute)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductSpecificationAttributeRepository creates a new instance of ProductSpecificationAttributeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductSpecificationAttributeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductSpecificationAttributeRepository {
	mock := &ProductSpecificationAttributeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
