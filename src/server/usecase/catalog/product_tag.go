package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type producttagUsecase struct {
	producttagRepository domain.ProductTagRepository
	contextTimeout       time.Duration
}

func NewProductTagUsecase(producttagRepository domain.ProductTagRepository, timeout time.Duration) domain.ProductTagUsecase {
	return &producttagUsecase{
		producttagRepository: producttagRepository,
		contextTimeout:       timeout,
	}
}

func (tu *producttagUsecase) CreateMany(c context.Context, items []domain.ProductTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttagRepository.CreateMany(ctx, items)
}

func (tu *producttagUsecase) Create(c context.Context, producttag *domain.ProductTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttagRepository.Create(ctx, producttag)
}

func (tu *producttagUsecase) Update(c context.Context, producttag *domain.ProductTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttagRepository.Update(ctx, producttag)
}

func (tu *producttagUsecase) Delete(c context.Context, producttag string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttagRepository.Delete(ctx, producttag)
}

func (lu *producttagUsecase) FetchByID(c context.Context, producttagID string) (domain.ProductTag, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.producttagRepository.FetchByID(ctx, producttagID)
}

func (lu *producttagUsecase) Fetch(c context.Context) ([]domain.ProductTag, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.producttagRepository.Fetch(ctx)
}
