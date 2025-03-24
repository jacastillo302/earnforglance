package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type adminareasettingsUsecase struct {
	adminareasettingsRepository domain.AdminAreaSettingsRepository
	contextTimeout              time.Duration
}

func NewAdminAreaSettingsUsecase(adminareasettingsRepository domain.AdminAreaSettingsRepository, timeout time.Duration) domain.AdminAreaSettingsUsecase {
	return &adminareasettingsUsecase{
		adminareasettingsRepository: adminareasettingsRepository,
		contextTimeout:              timeout,
	}
}

func (tu *adminareasettingsUsecase) CreateMany(c context.Context, items []domain.AdminAreaSettings) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.adminareasettingsRepository.CreateMany(ctx, items)
}

func (tu *adminareasettingsUsecase) Create(c context.Context, adminareasettings *domain.AdminAreaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.adminareasettingsRepository.Create(ctx, adminareasettings)
}

func (tu *adminareasettingsUsecase) Update(c context.Context, adminareasettings *domain.AdminAreaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.adminareasettingsRepository.Update(ctx, adminareasettings)
}

func (tu *adminareasettingsUsecase) Delete(c context.Context, adminareasettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.adminareasettingsRepository.Delete(ctx, adminareasettings)
}

func (lu *adminareasettingsUsecase) FetchByID(c context.Context, adminareasettingsID string) (domain.AdminAreaSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.adminareasettingsRepository.FetchByID(ctx, adminareasettingsID)
}

func (lu *adminareasettingsUsecase) Fetch(c context.Context) ([]domain.AdminAreaSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.adminareasettingsRepository.Fetch(ctx)
}
