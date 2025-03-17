package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/affiliate"
)

type affiliateUsecase struct {
	affiliateRepository domain.AffiliateRepository
	contextTimeout      time.Duration
}

func NewAffiliateUsecase(affiliateRepository domain.AffiliateRepository, timeout time.Duration) domain.AffiliateUsecase {
	return &affiliateUsecase{
		affiliateRepository: affiliateRepository,
		contextTimeout:      timeout,
	}
}

func (tu *affiliateUsecase) Create(c context.Context, affiliate *domain.Affiliate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.affiliateRepository.Create(ctx, affiliate)
}

func (tu *affiliateUsecase) Update(c context.Context, affiliate *domain.Affiliate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.affiliateRepository.Update(ctx, affiliate)
}

func (lu *affiliateUsecase) FetchByID(c context.Context, affiliateID string) (domain.Affiliate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.affiliateRepository.FetchByID(ctx, affiliateID)
}

func (lu *affiliateUsecase) Fetch(c context.Context) ([]domain.Affiliate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.affiliateRepository.Fetch(ctx)
}
