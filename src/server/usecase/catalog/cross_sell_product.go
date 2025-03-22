package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type crosssellproductUsecase struct {
	crosssellproductRepository domain.CrossSellProductRepository
	contextTimeout             time.Duration
}

func NewCrossSellProductUsecase(crosssellproductRepository domain.CrossSellProductRepository, timeout time.Duration) domain.CrossSellProductUsecase {
	return &crosssellproductUsecase{
		crosssellproductRepository: crosssellproductRepository,
		contextTimeout:             timeout,
	}
}

func (tu *crosssellproductUsecase) Create(c context.Context, affiliate *domain.CrossSellProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.Create(ctx, affiliate)
}

func (tu *crosssellproductUsecase) Update(c context.Context, affiliate *domain.CrossSellProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.Update(ctx, affiliate)
}

func (tu *crosssellproductUsecase) Delete(c context.Context, affiliate string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.Delete(ctx, affiliate)
}

func (lu *crosssellproductUsecase) FetchByID(c context.Context, affiliateID string) (domain.CrossSellProduct, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.crosssellproductRepository.FetchByID(ctx, affiliateID)
}

func (lu *crosssellproductUsecase) Fetch(c context.Context) ([]domain.CrossSellProduct, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.crosssellproductRepository.Fetch(ctx)
}
