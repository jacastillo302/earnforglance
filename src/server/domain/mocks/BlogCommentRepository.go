// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "earnforglance/server/domain/blogs"

	mock "github.com/stretchr/testify/mock"
)

// BlogCommentRepository is an autogenerated mock type for the BlogCommentRepository type
type BlogCommentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, blog_comment
func (_m *BlogCommentRepository) Create(c context.Context, blog_comment *domain.BlogComment) error {
	ret := _m.Called(c, blog_comment)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BlogComment) error); ok {
		r0 = rf(c, blog_comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: c, items
func (_m *BlogCommentRepository) CreateMany(c context.Context, items []domain.BlogComment) error {
	ret := _m.Called(c, items)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.BlogComment) error); ok {
		r0 = rf(c, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, ID
func (_m *BlogCommentRepository) Delete(c context.Context, ID string) error {
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
func (_m *BlogCommentRepository) Fetch(c context.Context) ([]domain.BlogComment, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.BlogComment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.BlogComment, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.BlogComment); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BlogComment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: c, blog_commentID
func (_m *BlogCommentRepository) FetchByID(c context.Context, blog_commentID string) (domain.BlogComment, error) {
	ret := _m.Called(c, blog_commentID)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 domain.BlogComment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.BlogComment, error)); ok {
		return rf(c, blog_commentID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.BlogComment); ok {
		r0 = rf(c, blog_commentID)
	} else {
		r0 = ret.Get(0).(domain.BlogComment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, blog_commentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, blog_comment
func (_m *BlogCommentRepository) Update(c context.Context, blog_comment *domain.BlogComment) error {
	ret := _m.Called(c, blog_comment)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BlogComment) error); ok {
		r0 = rf(c, blog_comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlogCommentRepository creates a new instance of BlogCommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogCommentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogCommentRepository {
	mock := &BlogCommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
