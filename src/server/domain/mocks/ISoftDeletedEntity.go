// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ISoftDeletedEntity is an autogenerated mock type for the ISoftDeletedEntity type
type ISoftDeletedEntity struct {
	mock.Mock
}

// GetDeleted provides a mock function with no fields
func (_m *ISoftDeletedEntity) GetDeleted() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDeleted")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetDeleted provides a mock function with given fields: deleted
func (_m *ISoftDeletedEntity) SetDeleted(deleted bool) {
	_m.Called(deleted)
}

// NewISoftDeletedEntity creates a new instance of ISoftDeletedEntity. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewISoftDeletedEntity(t interface {
	mock.TestingT
	Cleanup(func())
}) *ISoftDeletedEntity {
	mock := &ISoftDeletedEntity{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
