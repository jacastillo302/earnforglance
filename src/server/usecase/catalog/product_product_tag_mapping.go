package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type ProductProductTagMappingUsecase struct {
	ProductProductTagMappingRepository domain.ProductProductTagMappingRepository
	contextTimeout                     time.Duration
}

func NewProductProductTagMappingUsecase(ProductProductTagMappingRepository domain.ProductProductTagMappingRepository, timeout time.Duration) domain.ProductProductTagMappingUsecase {
	return &ProductProductTagMappingUsecase{
		ProductProductTagMappingRepository: ProductProductTagMappingRepository,
		contextTimeout:                     timeout,
	}
}

func (tu *ProductProductTagMappingUsecase) CreateMany(c context.Context, items []domain.ProductProductTagMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductProductTagMappingRepository.CreateMany(ctx, items)
}

func (tu *ProductProductTagMappingUsecase) Create(c context.Context, ProductProductTagMapping *domain.ProductProductTagMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductProductTagMappingRepository.Create(ctx, ProductProductTagMapping)
}

func (tu *ProductProductTagMappingUsecase) Update(c context.Context, ProductProductTagMapping *domain.ProductProductTagMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductProductTagMappingRepository.Update(ctx, ProductProductTagMapping)
}

func (tu *ProductProductTagMappingUsecase) Delete(c context.Context, ProductProductTagMapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductProductTagMappingRepository.Delete(ctx, ProductProductTagMapping)
}

func (lu *ProductProductTagMappingUsecase) FetchByID(c context.Context, ProductProductTagMappingID string) (domain.ProductProductTagMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.ProductProductTagMappingRepository.FetchByID(ctx, ProductProductTagMappingID)
}

func (lu *ProductProductTagMappingUsecase) Fetch(c context.Context) ([]domain.ProductProductTagMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.ProductProductTagMappingRepository.Fetch(ctx)
}
