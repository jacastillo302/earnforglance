package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountrequirementUsecase struct {
	discountrequirementRepository domain.DiscountRequirementRepository
	contextTimeout                time.Duration
}

func NewDiscountRequirementUsecase(discountrequirementRepository domain.DiscountRequirementRepository, timeout time.Duration) domain.DiscountRequirementUsecase {
	return &discountrequirementUsecase{
		discountrequirementRepository: discountrequirementRepository,
		contextTimeout:                timeout,
	}
}

func (tu *discountrequirementUsecase) Create(c context.Context, discountrequirement *domain.DiscountRequirement) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountrequirementRepository.Create(ctx, discountrequirement)
}

func (tu *discountrequirementUsecase) Update(c context.Context, discountrequirement *domain.DiscountRequirement) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountrequirementRepository.Update(ctx, discountrequirement)
}

func (tu *discountrequirementUsecase) Delete(c context.Context, discountrequirement string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountrequirementRepository.Delete(ctx, discountrequirement)
}

func (lu *discountrequirementUsecase) FetchByID(c context.Context, discountrequirementID string) (domain.DiscountRequirement, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountrequirementRepository.FetchByID(ctx, discountrequirementID)
}

func (lu *discountrequirementUsecase) Fetch(c context.Context) ([]domain.DiscountRequirement, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountrequirementRepository.Fetch(ctx)
}
