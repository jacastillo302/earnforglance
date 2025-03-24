package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type recurringpaymenthistoryUsecase struct {
	recurringpaymenthistoryRepository domain.RecurringPaymentHistoryRepository
	contextTimeout                    time.Duration
}

func NewRecurringPaymentHistoryUsecase(recurringpaymenthistoryRepository domain.RecurringPaymentHistoryRepository, timeout time.Duration) domain.RecurringPaymentHistoryUsecase {
	return &recurringpaymenthistoryUsecase{
		recurringpaymenthistoryRepository: recurringpaymenthistoryRepository,
		contextTimeout:                    timeout,
	}
}

func (tu *recurringpaymenthistoryUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.RecurringPaymentHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymenthistoryRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *recurringpaymenthistoryUsecase) Create(c context.Context, recurringpaymenthistory *domain.RecurringPaymentHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymenthistoryRepository.Create(ctx, recurringpaymenthistory)
}

func (tu *recurringpaymenthistoryUsecase) Update(c context.Context, recurringpaymenthistory *domain.RecurringPaymentHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymenthistoryRepository.Update(ctx, recurringpaymenthistory)
}

func (tu *recurringpaymenthistoryUsecase) Delete(c context.Context, recurringpaymenthistory string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.recurringpaymenthistoryRepository.Delete(ctx, recurringpaymenthistory)
}

func (lu *recurringpaymenthistoryUsecase) FetchByID(c context.Context, recurringpaymenthistoryID string) (domain.RecurringPaymentHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.recurringpaymenthistoryRepository.FetchByID(ctx, recurringpaymenthistoryID)
}

func (lu *recurringpaymenthistoryUsecase) Fetch(c context.Context) ([]domain.RecurringPaymentHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.recurringpaymenthistoryRepository.Fetch(ctx)
}
