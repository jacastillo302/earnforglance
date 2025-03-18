package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productManufacturerUsecase struct {
	productManufacturerRepository domain.ProductManufacturerRepository
	contextTimeout                time.Duration
}

func NewProductManufacturerUsecase(productManufacturerRepository domain.ProductManufacturerRepository, timeout time.Duration) domain.ProductManufacturerUsecase {
	return &productManufacturerUsecase{
		productManufacturerRepository: productManufacturerRepository,
		contextTimeout:                timeout,
	}
}

func (tu *productManufacturerUsecase) Create(c context.Context, productManufacturer *domain.ProductManufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productManufacturerRepository.Create(ctx, productManufacturer)
}

func (tu *productManufacturerUsecase) Update(c context.Context, productManufacturer *domain.ProductManufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productManufacturerRepository.Update(ctx, productManufacturer)
}

func (tu *productManufacturerUsecase) Delete(c context.Context, productManufacturer *domain.ProductManufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productManufacturerRepository.Delete(ctx, productManufacturer)
}

func (lu *productManufacturerUsecase) FetchByID(c context.Context, productManufacturerID string) (domain.ProductManufacturer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productManufacturerRepository.FetchByID(ctx, productManufacturerID)
}

func (lu *productManufacturerUsecase) Fetch(c context.Context) ([]domain.ProductManufacturer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productManufacturerRepository.Fetch(ctx)
}
