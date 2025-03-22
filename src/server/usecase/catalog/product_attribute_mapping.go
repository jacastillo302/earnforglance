package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productattributemappingUsecase struct {
	productattributemappingRepository domain.ProductAttributeMappingRepository
	contextTimeout                    time.Duration
}

func NewProductAttributeMappingUsecase(productattributemappingRepository domain.ProductAttributeMappingRepository, timeout time.Duration) domain.ProductAttributeMappingUsecase {
	return &productattributemappingUsecase{
		productattributemappingRepository: productattributemappingRepository,
		contextTimeout:                    timeout,
	}
}

func (tu *productattributemappingUsecase) Create(c context.Context, productattributemapping *domain.ProductAttributeMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributemappingRepository.Create(ctx, productattributemapping)
}

func (tu *productattributemappingUsecase) Update(c context.Context, productattributemapping *domain.ProductAttributeMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributemappingRepository.Update(ctx, productattributemapping)
}

func (tu *productattributemappingUsecase) Delete(c context.Context, productattributemapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributemappingRepository.Delete(ctx, productattributemapping)
}

func (lu *productattributemappingUsecase) FetchByID(c context.Context, productattributemappingID string) (domain.ProductAttributeMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributemappingRepository.FetchByID(ctx, productattributemappingID)
}

func (lu *productattributemappingUsecase) Fetch(c context.Context) ([]domain.ProductAttributeMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributemappingRepository.Fetch(ctx)
}
