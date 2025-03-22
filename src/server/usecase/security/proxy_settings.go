package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type proxySettingsUsecase struct {
	proxySettingsRepository domain.ProxySettingsRepository
	contextTimeout          time.Duration
}

func NewProxySettingsUsecase(proxySettingsRepository domain.ProxySettingsRepository, timeout time.Duration) domain.ProxySettingsUsecase {
	return &proxySettingsUsecase{
		proxySettingsRepository: proxySettingsRepository,
		contextTimeout:          timeout,
	}
}

func (tu *proxySettingsUsecase) Create(c context.Context, proxySettings *domain.ProxySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.proxySettingsRepository.Create(ctx, proxySettings)
}

func (tu *proxySettingsUsecase) Update(c context.Context, proxySettings *domain.ProxySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.proxySettingsRepository.Update(ctx, proxySettings)
}

func (tu *proxySettingsUsecase) Delete(c context.Context, proxySettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.proxySettingsRepository.Delete(ctx, proxySettings)
}

func (lu *proxySettingsUsecase) FetchByID(c context.Context, proxySettingsID string) (domain.ProxySettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.proxySettingsRepository.FetchByID(ctx, proxySettingsID)
}

func (lu *proxySettingsUsecase) Fetch(c context.Context) ([]domain.ProxySettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.proxySettingsRepository.Fetch(ctx)
}
