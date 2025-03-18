package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productattributecombinationUsecase struct {
	productattributecombinationRepository domain.ProductAttributeCombinationRepository
	contextTimeout                        time.Duration
}

func NewProductAttributeCombinationUsecase(productattributecombinationRepository domain.ProductAttributeCombinationRepository, timeout time.Duration) domain.ProductAttributeCombinationUsecase {
	return &productattributecombinationUsecase{
		productattributecombinationRepository: productattributecombinationRepository,
		contextTimeout:                        timeout,
	}
}

func (tu *productattributecombinationUsecase) Create(c context.Context, productattributecombination *domain.ProductAttributeCombination) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributecombinationRepository.Create(ctx, productattributecombination)
}

func (tu *productattributecombinationUsecase) Update(c context.Context, productattributecombination *domain.ProductAttributeCombination) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributecombinationRepository.Update(ctx, productattributecombination)
}

func (tu *productattributecombinationUsecase) Delete(c context.Context, productattributecombination *domain.ProductAttributeCombination) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributecombinationRepository.Delete(ctx, productattributecombination)
}

func (lu *productattributecombinationUsecase) FetchByID(c context.Context, productattributecombinationID string) (domain.ProductAttributeCombination, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributecombinationRepository.FetchByID(ctx, productattributecombinationID)
}

func (lu *productattributecombinationUsecase) Fetch(c context.Context) ([]domain.ProductAttributeCombination, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributecombinationRepository.Fetch(ctx)
}
