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

func (tu *crosssellproductUsecase) CreateMany(c context.Context, items []domain.CrossSellProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.CreateMany(ctx, items)
}

func (tu *crosssellproductUsecase) Create(c context.Context, item *domain.CrossSellProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.Create(ctx, item)
}

func (tu *crosssellproductUsecase) Update(c context.Context, item *domain.CrossSellProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.Update(ctx, item)
}

func (tu *crosssellproductUsecase) Delete(c context.Context, item string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.crosssellproductRepository.Delete(ctx, item)
}

func (lu *crosssellproductUsecase) FetchByID(c context.Context, itemID string) (domain.CrossSellProduct, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.crosssellproductRepository.FetchByID(ctx, itemID)
}

func (lu *crosssellproductUsecase) Fetch(c context.Context) ([]domain.CrossSellProduct, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.crosssellproductRepository.Fetch(ctx)
}
