package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountproductmappingUsecase struct {
	discountproductmappingRepository domain.DiscountProductMappingRepository
	contextTimeout                   time.Duration
}

func NewDiscountProductMappingUsecase(discountproductmappingRepository domain.DiscountProductMappingRepository, timeout time.Duration) domain.DiscountProductMappingUsecase {
	return &discountproductmappingUsecase{
		discountproductmappingRepository: discountproductmappingRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *discountproductmappingUsecase) Create(c context.Context, discountproductmapping *domain.DiscountProductMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountproductmappingRepository.Create(ctx, discountproductmapping)
}

func (tu *discountproductmappingUsecase) Update(c context.Context, discountproductmapping *domain.DiscountProductMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountproductmappingRepository.Update(ctx, discountproductmapping)
}

func (tu *discountproductmappingUsecase) Delete(c context.Context, discountproductmapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountproductmappingRepository.Delete(ctx, discountproductmapping)
}

func (lu *discountproductmappingUsecase) FetchByID(c context.Context, discountproductmappingID string) (domain.DiscountProductMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountproductmappingRepository.FetchByID(ctx, discountproductmappingID)
}

func (lu *discountproductmappingUsecase) Fetch(c context.Context) ([]domain.DiscountProductMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountproductmappingRepository.Fetch(ctx)
}
