package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type blogUsecase struct {
	itemRepository domain.BlogRepository
	contextTimeout time.Duration
}

func NewblogUsecase(itemRepository domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (r *blogUsecase) GetBlogs(c context.Context, filter domain.BlogRequest) ([]domain.BlogsResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.itemRepository.GetBlogs(ctx, filter)
}
