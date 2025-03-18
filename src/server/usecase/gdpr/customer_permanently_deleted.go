package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/gdpr"
)

type customerPermanentlyDeletedUsecase struct {
	customerPermanentlyDeletedRepository domain.CustomerPermanentlyDeletedRepository
	contextTimeout                       time.Duration
}

func NewCustomerPermanentlyDeletedUsecase(customerPermanentlyDeletedRepository domain.CustomerPermanentlyDeletedRepository, timeout time.Duration) domain.CustomerPermanentlyDeletedUsecase {
	return &customerPermanentlyDeletedUsecase{
		customerPermanentlyDeletedRepository: customerPermanentlyDeletedRepository,
		contextTimeout:                       timeout,
	}
}

func (tu *customerPermanentlyDeletedUsecase) Create(c context.Context, customerPermanentlyDeleted *domain.CustomerPermanentlyDeleted) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerPermanentlyDeletedRepository.Create(ctx, customerPermanentlyDeleted)
}

func (tu *customerPermanentlyDeletedUsecase) Update(c context.Context, customerPermanentlyDeleted *domain.CustomerPermanentlyDeleted) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerPermanentlyDeletedRepository.Update(ctx, customerPermanentlyDeleted)
}

func (tu *customerPermanentlyDeletedUsecase) Delete(c context.Context, customerPermanentlyDeleted *domain.CustomerPermanentlyDeleted) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerPermanentlyDeletedRepository.Delete(ctx, customerPermanentlyDeleted)
}

func (lu *customerPermanentlyDeletedUsecase) FetchByID(c context.Context, customerPermanentlyDeletedID string) (domain.CustomerPermanentlyDeleted, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerPermanentlyDeletedRepository.FetchByID(ctx, customerPermanentlyDeletedID)
}

func (lu *customerPermanentlyDeletedUsecase) Fetch(c context.Context) ([]domain.CustomerPermanentlyDeleted, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerPermanentlyDeletedRepository.Fetch(ctx)
}
