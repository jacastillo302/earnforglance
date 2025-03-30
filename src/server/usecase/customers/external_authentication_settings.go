package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type externalAuthenticationSettingsUsecase struct {
	externalAuthenticationSettingsRepository domain.ExternalAuthenticationSettingsRepository
	contextTimeout                           time.Duration
}

func NewExternalAuthenticationSettingsUsecase(externalAuthenticationSettingsRepository domain.ExternalAuthenticationSettingsRepository, timeout time.Duration) domain.ExternalAuthenticationSettingsUsecase {
	return &externalAuthenticationSettingsUsecase{
		externalAuthenticationSettingsRepository: externalAuthenticationSettingsRepository,
		contextTimeout:                           timeout,
	}
}

func (tu *externalAuthenticationSettingsUsecase) CreateMany(c context.Context, items []domain.ExternalAuthenticationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationSettingsRepository.CreateMany(ctx, items)
}

func (tu *externalAuthenticationSettingsUsecase) Create(c context.Context, externalAuthenticationSettings *domain.ExternalAuthenticationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationSettingsRepository.Create(ctx, externalAuthenticationSettings)
}

func (tu *externalAuthenticationSettingsUsecase) Update(c context.Context, externalAuthenticationSettings *domain.ExternalAuthenticationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationSettingsRepository.Update(ctx, externalAuthenticationSettings)
}

func (tu *externalAuthenticationSettingsUsecase) Delete(c context.Context, externalAuthenticationSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.externalAuthenticationSettingsRepository.Delete(ctx, externalAuthenticationSettings)
}

func (lu *externalAuthenticationSettingsUsecase) FetchByID(c context.Context, externalAuthenticationSettingsID string) (domain.ExternalAuthenticationSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.externalAuthenticationSettingsRepository.FetchByID(ctx, externalAuthenticationSettingsID)
}

func (lu *externalAuthenticationSettingsUsecase) Fetch(c context.Context) ([]domain.ExternalAuthenticationSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.externalAuthenticationSettingsRepository.Fetch(ctx)
}
