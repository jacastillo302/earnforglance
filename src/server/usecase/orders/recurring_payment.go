package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type recurringpaymentUsecase struct {
	recurringpaymentRepository domain.RecurringPaymentRepository
	contextTimeout             time.Duration
}

func NewRecurringPaymentUsecase(recurringpaymentRepository domain.RecurringPaymentRepository, timeout time.Duration) domain.RecurringPaymentUsecase {
	return &recurringpaymentUsecase{
		recurringpaymentRepository: recurringpaymentRepository,
		contextTimeout:             timeout,
	}
}

func (tu *recurringpaymentUsecase) Create(c context.Context, recurringpayment *domain.RecurringPayment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymentRepository.Create(ctx, recurringpayment)
}

func (tu *recurringpaymentUsecase) Update(c context.Context, recurringpayment *domain.RecurringPayment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymentRepository.Update(ctx, recurringpayment)
}

func (tu *recurringpaymentUsecase) Delete(c context.Context, recurringpayment string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymentRepository.Delete(ctx, recurringpayment)
}

func (lu *recurringpaymentUsecase) FetchByID(c context.Context, recurringpaymentID string) (domain.RecurringPayment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.recurringpaymentRepository.FetchByID(ctx, recurringpaymentID)
}

func (lu *recurringpaymentUsecase) Fetch(c context.Context) ([]domain.RecurringPayment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.recurringpaymentRepository.Fetch(ctx)
}
