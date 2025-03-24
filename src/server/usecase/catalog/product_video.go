package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productvideoUsecase struct {
	productvideoRepository domain.ProductVideoRepository
	contextTimeout         time.Duration
}

func NewProductVideoUsecase(productvideoRepository domain.ProductVideoRepository, timeout time.Duration) domain.ProductVideoUsecase {
	return &productvideoUsecase{
		productvideoRepository: productvideoRepository,
		contextTimeout:         timeout,
	}
}

func (tu *productvideoUsecase) CreateMany(c context.Context, items []domain.ProductVideo) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productvideoRepository.CreateMany(ctx, items)
}

func (tu *productvideoUsecase) Create(c context.Context, productvideo *domain.ProductVideo) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productvideoRepository.Create(ctx, productvideo)
}

func (tu *productvideoUsecase) Update(c context.Context, productvideo *domain.ProductVideo) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productvideoRepository.Update(ctx, productvideo)
}

func (tu *productvideoUsecase) Delete(c context.Context, productvideo string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productvideoRepository.Delete(ctx, productvideo)
}

func (lu *productvideoUsecase) FetchByID(c context.Context, productvideoID string) (domain.ProductVideo, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productvideoRepository.FetchByID(ctx, productvideoID)
}

func (lu *productvideoUsecase) Fetch(c context.Context) ([]domain.ProductVideo, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productvideoRepository.Fetch(ctx)
}
