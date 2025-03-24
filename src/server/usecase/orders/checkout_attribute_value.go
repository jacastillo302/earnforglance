package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

// checkoutattributevalueUsecase represents the usecase for the checkout attribute value
type checkoutattributevalueUsecase struct {
	checkoutattributevalueRepository domain.CheckoutAttributeValueRepository
	contextTimeout                   time.Duration
}

func NewCheckoutAttributeValueUsecase(checkoutattributevalueRepository domain.CheckoutAttributeValueRepository, timeout time.Duration) domain.CheckoutAttributeValueUsecase {
	return &checkoutattributevalueUsecase{
		checkoutattributevalueRepository: checkoutattributevalueRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *checkoutattributevalueUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.CheckoutAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributevalueRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *checkoutattributevalueUsecase) Create(c context.Context, checkoutattributevalue *domain.CheckoutAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributevalueRepository.Create(ctx, checkoutattributevalue)
}

func (tu *checkoutattributevalueUsecase) Update(c context.Context, checkoutattributevalue *domain.CheckoutAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributevalueRepository.Update(ctx, checkoutattributevalue)
}

func (tu *checkoutattributevalueUsecase) Delete(c context.Context, checkoutattributevalue string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.checkoutattributevalueRepository.Delete(ctx, checkoutattributevalue)
}

func (lu *checkoutattributevalueUsecase) FetchByID(c context.Context, checkoutattributevalueID string) (domain.CheckoutAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.checkoutattributevalueRepository.FetchByID(ctx, checkoutattributevalueID)
}

func (lu *checkoutattributevalueUsecase) Fetch(c context.Context) ([]domain.CheckoutAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.checkoutattributevalueRepository.Fetch(ctx)
}
