package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productattributevalueUsecase struct {
	productattributevalueRepository domain.ProductAttributeValueRepository
	contextTimeout                  time.Duration
}

func NewProductAttributeValueUsecase(productattributevalueRepository domain.ProductAttributeValueRepository, timeout time.Duration) domain.ProductAttributeValueUsecase {
	return &productattributevalueUsecase{
		productattributevalueRepository: productattributevalueRepository,
		contextTimeout:                  timeout,
	}
}

func (tu *productattributevalueUsecase) Create(c context.Context, productattributevalue *domain.ProductAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributevalueRepository.Create(ctx, productattributevalue)
}

func (tu *productattributevalueUsecase) Update(c context.Context, productattributevalue *domain.ProductAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributevalueRepository.Update(ctx, productattributevalue)
}

func (tu *productattributevalueUsecase) Delete(c context.Context, productattributevalue string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributevalueRepository.Delete(ctx, productattributevalue)
}

func (lu *productattributevalueUsecase) FetchByID(c context.Context, productattributevalueID string) (domain.ProductAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributevalueRepository.FetchByID(ctx, productattributevalueID)
}

func (lu *productattributevalueUsecase) Fetch(c context.Context) ([]domain.ProductAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributevalueRepository.Fetch(ctx)
}
