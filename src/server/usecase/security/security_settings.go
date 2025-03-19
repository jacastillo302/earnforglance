package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type securitySettingsUsecase struct {
	securitySettingsRepository domain.SecuritySettingsRepository
	contextTimeout             time.Duration
}

func NewSecuritySettingsUsecase(securitySettingsRepository domain.SecuritySettingsRepository, timeout time.Duration) domain.SecuritySettingsUsecase {
	return &securitySettingsUsecase{
		securitySettingsRepository: securitySettingsRepository,
		contextTimeout:             timeout,
	}
}

func (tu *securitySettingsUsecase) Create(c context.Context, securitySettings *domain.SecuritySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.securitySettingsRepository.Create(ctx, securitySettings)
}

func (tu *securitySettingsUsecase) Update(c context.Context, securitySettings *domain.SecuritySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.securitySettingsRepository.Update(ctx, securitySettings)
}

func (tu *securitySettingsUsecase) Delete(c context.Context, securitySettings *domain.SecuritySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.securitySettingsRepository.Delete(ctx, securitySettings)
}

func (lu *securitySettingsUsecase) FetchByID(c context.Context, securitySettingsID string) (domain.SecuritySettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.securitySettingsRepository.FetchByID(ctx, securitySettingsID)
}

func (lu *securitySettingsUsecase) Fetch(c context.Context) ([]domain.SecuritySettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.securitySettingsRepository.Fetch(ctx)
}
