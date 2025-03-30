package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type multiFactorAuthenticationSettingsUsecase struct {
	multiFactorAuthenticationSettingsRepository domain.MultiFactorAuthenticationSettingsRepository
	contextTimeout                              time.Duration
}

func NewMultiFactorAuthenticationSettingsUsecase(multiFactorAuthenticationSettingsRepository domain.MultiFactorAuthenticationSettingsRepository, timeout time.Duration) domain.MultiFactorAuthenticationSettingsUsecase {
	return &multiFactorAuthenticationSettingsUsecase{
		multiFactorAuthenticationSettingsRepository: multiFactorAuthenticationSettingsRepository,
		contextTimeout: timeout,
	}
}

func (tu *multiFactorAuthenticationSettingsUsecase) CreateMany(c context.Context, items []domain.MultiFactorAuthenticationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.multiFactorAuthenticationSettingsRepository.CreateMany(ctx, items)
}

func (tu *multiFactorAuthenticationSettingsUsecase) Create(c context.Context, multiFactorAuthenticationSettings *domain.MultiFactorAuthenticationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.multiFactorAuthenticationSettingsRepository.Create(ctx, multiFactorAuthenticationSettings)
}

func (tu *multiFactorAuthenticationSettingsUsecase) Update(c context.Context, multiFactorAuthenticationSettings *domain.MultiFactorAuthenticationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.multiFactorAuthenticationSettingsRepository.Update(ctx, multiFactorAuthenticationSettings)
}

func (tu *multiFactorAuthenticationSettingsUsecase) Delete(c context.Context, multiFactorAuthenticationSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.multiFactorAuthenticationSettingsRepository.Delete(ctx, multiFactorAuthenticationSettings)
}

func (lu *multiFactorAuthenticationSettingsUsecase) FetchByID(c context.Context, multiFactorAuthenticationSettingsID string) (domain.MultiFactorAuthenticationSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.multiFactorAuthenticationSettingsRepository.FetchByID(ctx, multiFactorAuthenticationSettingsID)
}

func (lu *multiFactorAuthenticationSettingsUsecase) Fetch(c context.Context) ([]domain.MultiFactorAuthenticationSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.multiFactorAuthenticationSettingsRepository.Fetch(ctx)
}
