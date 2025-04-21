package domain

import (
	"context"
	domain "earnforglance/server/domain/blogs"
)

type BlogRequest struct {
	ID            string
	Filters       []Filter
	Sort          string
	Limit         int
	Page          int
	Lang          string
	AllowComments bool
	Content       []string
}

type BlogResponse struct {
	Blog    domain.BlogPost
	Coments []domain.BlogComment
	Tags    []domain.BlogPostTag
}

type BlogsResponse struct {
	Blogs []BlogResponse
}

type BlogRepository interface {
	GetBlogs(c context.Context, filter BlogRequest) ([]BlogsResponse, error)
}

type BlogUsecase interface {
	GetBlogs(c context.Context, filter BlogRequest) ([]BlogsResponse, error)
}
