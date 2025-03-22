package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productattributeUsecase struct {
	productattributeRepository domain.ProductAttributeRepository
	contextTimeout             time.Duration
}

func NewProductAttributeUsecase(productattributeRepository domain.ProductAttributeRepository, timeout time.Duration) domain.ProductAttributeUsecase {
	return &productattributeUsecase{
		productattributeRepository: productattributeRepository,
		contextTimeout:             timeout,
	}
}

func (tu *productattributeUsecase) Create(c context.Context, productattribute *domain.ProductAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributeRepository.Create(ctx, productattribute)
}

func (tu *productattributeUsecase) Update(c context.Context, productattribute *domain.ProductAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributeRepository.Update(ctx, productattribute)
}

func (tu *productattributeUsecase) Delete(c context.Context, productattribute string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributeRepository.Delete(ctx, productattribute)
}

func (lu *productattributeUsecase) FetchByID(c context.Context, productattributeID string) (domain.ProductAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributeRepository.FetchByID(ctx, productattributeID)
}

func (lu *productattributeUsecase) Fetch(c context.Context) ([]domain.ProductAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributeRepository.Fetch(ctx)
}
