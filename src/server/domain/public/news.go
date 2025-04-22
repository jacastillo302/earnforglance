package domain

import (
	"context"
	domain "earnforglance/server/domain/news"
)

type NewsItemRequest struct {
	ID            string
	Filters       []Filter
	Sort          string
	Limit         int
	Page          int
	Lang          string
	AllowComments bool
	Content       []string
}

type NewsItemResponse struct {
	News     domain.NewsItem
	Comments []domain.NewsComment
}

type NewsItemsResponse struct {
	News []NewsItemResponse
}

type NewsItemRepository interface {
	GetNewsItems(c context.Context, filter NewsItemRequest) ([]NewsItemsResponse, error)
}

type NewsItemUsecase interface {
	GetNewsItems(c context.Context, filter NewsItemRequest) ([]NewsItemsResponse, error)
}
