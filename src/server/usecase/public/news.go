package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type newsItemUsecase struct {
	itemRepository domain.NewsItemRepository
	contextTimeout time.Duration
}

func NewnewsItemUsecase(itemRepository domain.NewsItemRepository, timeout time.Duration) domain.NewsItemUsecase {
	return &newsItemUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (r *newsItemUsecase) GetNewsItems(c context.Context, filter domain.NewsItemRequest) ([]domain.NewsItemsResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.itemRepository.GetNewsItems(ctx, filter)
}
