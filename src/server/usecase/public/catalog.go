package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type catalogtUsecase struct {
	itemRepository domain.CatalogRepository
	contextTimeout time.Duration
}

func NewCatalogtUsecase(itemRepository domain.CatalogRepository, timeout time.Duration) domain.CatalogtUsecase {
	return &catalogtUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *catalogtUsecase) GetProduct(c context.Context, ID string) (domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetProduct(ctx, ID)
}
