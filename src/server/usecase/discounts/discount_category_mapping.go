package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountcategorymappingUsecase struct {
	discountcategorymappingRepository domain.DiscountCategoryMappingRepository
	contextTimeout                    time.Duration
}

func NewDiscountCategoryMappingUsecase(discountcategorymappingRepository domain.DiscountCategoryMappingRepository, timeout time.Duration) domain.DiscountCategoryMappingUsecase {
	return &discountcategorymappingUsecase{
		discountcategorymappingRepository: discountcategorymappingRepository,
		contextTimeout:                    timeout,
	}
}

func (tu *discountcategorymappingUsecase) CreateMany(c context.Context, items []domain.DiscountCategoryMapping) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountcategorymappingRepository.CreateMany(ctx, items)
}

func (tu *discountcategorymappingUsecase) Create(c context.Context, discountcategorymapping *domain.DiscountCategoryMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountcategorymappingRepository.Create(ctx, discountcategorymapping)
}

func (tu *discountcategorymappingUsecase) Update(c context.Context, discountcategorymapping *domain.DiscountCategoryMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountcategorymappingRepository.Update(ctx, discountcategorymapping)
}

func (tu *discountcategorymappingUsecase) Delete(c context.Context, discountcategorymapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountcategorymappingRepository.Delete(ctx, discountcategorymapping)
}

func (lu *discountcategorymappingUsecase) FetchByID(c context.Context, discountcategorymappingID string) (domain.DiscountCategoryMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountcategorymappingRepository.FetchByID(ctx, discountcategorymappingID)
}

func (lu *discountcategorymappingUsecase) Fetch(c context.Context) ([]domain.DiscountCategoryMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountcategorymappingRepository.Fetch(ctx)
}
