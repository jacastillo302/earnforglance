package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productSpecificationAttributeUsecase struct {
	productSpecificationAttributeRepository domain.ProductSpecificationAttributeRepository
	contextTimeout                          time.Duration
}

func NewProductSpecificationAttributeUsecase(productSpecificationAttributeRepository domain.ProductSpecificationAttributeRepository, timeout time.Duration) domain.ProductSpecificationAttributeUsecase {
	return &productSpecificationAttributeUsecase{
		productSpecificationAttributeRepository: productSpecificationAttributeRepository,
		contextTimeout:                          timeout,
	}
}

func (tu *productSpecificationAttributeUsecase) Create(c context.Context, productSpecificationAttribute *domain.ProductSpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productSpecificationAttributeRepository.Create(ctx, productSpecificationAttribute)
}

func (tu *productSpecificationAttributeUsecase) Update(c context.Context, productSpecificationAttribute *domain.ProductSpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productSpecificationAttributeRepository.Update(ctx, productSpecificationAttribute)
}

func (tu *productSpecificationAttributeUsecase) Delete(c context.Context, productSpecificationAttribute *domain.ProductSpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productSpecificationAttributeRepository.Delete(ctx, productSpecificationAttribute)
}

func (lu *productSpecificationAttributeUsecase) FetchByID(c context.Context, productSpecificationAttributeID string) (domain.ProductSpecificationAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productSpecificationAttributeRepository.FetchByID(ctx, productSpecificationAttributeID)
}

func (lu *productSpecificationAttributeUsecase) Fetch(c context.Context) ([]domain.ProductSpecificationAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productSpecificationAttributeRepository.Fetch(ctx)
}
