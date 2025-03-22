package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type checkoutattributeUsecase struct {
	checkoutattributeRepository domain.CheckoutAttributeRepository
	contextTimeout              time.Duration
}

func NewCheckoutAttributeUsecase(checkoutattributeRepository domain.CheckoutAttributeRepository, timeout time.Duration) domain.CheckoutAttributeUsecase {
	return &checkoutattributeUsecase{
		checkoutattributeRepository: checkoutattributeRepository,
		contextTimeout:              timeout,
	}
}

func (tu *checkoutattributeUsecase) Create(c context.Context, checkoutattribute *domain.CheckoutAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributeRepository.Create(ctx, checkoutattribute)
}

func (tu *checkoutattributeUsecase) Update(c context.Context, checkoutattribute *domain.CheckoutAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributeRepository.Update(ctx, checkoutattribute)
}

func (tu *checkoutattributeUsecase) Delete(c context.Context, checkoutattribute string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributeRepository.Delete(ctx, checkoutattribute)
}

func (lu *checkoutattributeUsecase) FetchByID(c context.Context, checkoutattributeID string) (domain.CheckoutAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.checkoutattributeRepository.FetchByID(ctx, checkoutattributeID)
}

func (lu *checkoutattributeUsecase) Fetch(c context.Context) ([]domain.CheckoutAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.checkoutattributeRepository.Fetch(ctx)
}
