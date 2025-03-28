package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/api"
)

type apisettingsUsecase struct {
	apisettingsRepository domain.ApiSettingsRepository
	contextTimeout        time.Duration
}

func NewApiSettingsUsecase(apisettingsRepository domain.ApiSettingsRepository, timeout time.Duration) domain.ApiSettingsUsecase {
	return &apisettingsUsecase{
		apisettingsRepository: apisettingsRepository,
		contextTimeout:        timeout,
	}
}

func (tu *apisettingsUsecase) CreateMany(c context.Context, items []domain.ApiSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apisettingsRepository.CreateMany(ctx, items)
}

func (tu *apisettingsUsecase) Create(c context.Context, apisettings *domain.ApiSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apisettingsRepository.Create(ctx, apisettings)
}

func (tu *apisettingsUsecase) Update(c context.Context, apisettings *domain.ApiSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apisettingsRepository.Update(ctx, apisettings)
}

func (tu *apisettingsUsecase) Delete(c context.Context, apisettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apisettingsRepository.Delete(ctx, apisettings)
}

func (lu *apisettingsUsecase) FetchByID(c context.Context, apisettingsID string) (domain.ApiSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.apisettingsRepository.FetchByID(ctx, apisettingsID)
}

func (lu *apisettingsUsecase) Fetch(c context.Context) ([]domain.ApiSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.apisettingsRepository.Fetch(ctx)
}
