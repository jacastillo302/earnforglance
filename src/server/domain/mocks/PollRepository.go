// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/polls"

	mock "github.com/stretchr/testify/mock"
)

// PollRepository is an autogenerated mock type for the PollRepository type
type PollRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, poll
func (_m *PollRepository) Create(c context.Context, poll *domain.Poll) error {
	ret := _m.Called(c, poll)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Poll) error); ok {
		r0 = rf(c, poll)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *PollRepository) CreateMany(c context.Context, items []domain.Poll) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Poll) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *PollRepository) Delete(c context.Context, ID string) error {
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
func (_m *PollRepository) Fetch(c context.Context) ([]domain.Poll, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.Poll
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Poll, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Poll); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Poll)
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
func (_m *PollRepository) FetchByID(c context.Context, ID string) (domain.Poll, error) {
	ret := _m.Called(c, ID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.Poll
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Poll, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Poll); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(domain.Poll)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, poll
func (_m *PollRepository) Update(c context.Context, poll *domain.Poll) error {
	ret := _m.Called(c, poll)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Poll) error); ok {
		r0 = rf(c, poll)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPollRepository creates a new instance of PollRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPollRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PollRepository {
	mock := &PollRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
