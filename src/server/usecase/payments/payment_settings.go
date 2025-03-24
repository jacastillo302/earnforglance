package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/payments"
)

type paymentsettingsUsecase struct {
	paymentsettingsRepository domain.PaymentSettingsRepository
	contextTimeout            time.Duration
}

func NewPaymentSettingsUsecase(paymentsettingsRepository domain.PaymentSettingsRepository, timeout time.Duration) domain.PaymentSettingsUsecase {
	return &paymentsettingsUsecase{
		paymentsettingsRepository: paymentsettingsRepository,
		contextTimeout:            timeout,
	}
}

func (tu *paymentsettingsUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.PaymentSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.paymentsettingsRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *paymentsettingsUsecase) Create(c context.Context, paymentsettings *domain.PaymentSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.paymentsettingsRepository.Create(ctx, paymentsettings)
}

func (tu *paymentsettingsUsecase) Update(c context.Context, paymentsettings *domain.PaymentSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.paymentsettingsRepository.Update(ctx, paymentsettings)
}

func (tu *paymentsettingsUsecase) Delete(c context.Context, paymentsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.paymentsettingsRepository.Delete(ctx, paymentsettings)
}

func (lu *paymentsettingsUsecase) FetchByID(c context.Context, paymentsettingsID string) (domain.PaymentSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.paymentsettingsRepository.FetchByID(ctx, paymentsettingsID)
}

func (lu *paymentsettingsUsecase) Fetch(c context.Context) ([]domain.PaymentSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.paymentsettingsRepository.Fetch(ctx)
}
