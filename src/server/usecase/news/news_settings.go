package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/news"
)

type newssettingsUsecase struct {
	newssettingsRepository domain.NewsSettingsRepository
	contextTimeout         time.Duration
}

func NewNewsSettingsUsecase(newssettingsRepository domain.NewsSettingsRepository, timeout time.Duration) domain.NewsSettingsUsecase {
	return &newssettingsUsecase{
		newssettingsRepository: newssettingsRepository,
		contextTimeout:         timeout,
	}
}

func (tu *newssettingsUsecase) Create(c context.Context, newssettings *domain.NewsSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newssettingsRepository.Create(ctx, newssettings)
}

func (tu *newssettingsUsecase) Update(c context.Context, newssettings *domain.NewsSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newssettingsRepository.Update(ctx, newssettings)
}

func (tu *newssettingsUsecase) Delete(c context.Context, newssettings *domain.NewsSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newssettingsRepository.Delete(ctx, newssettings)
}

func (lu *newssettingsUsecase) FetchByID(c context.Context, newssettingsID string) (domain.NewsSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newssettingsRepository.FetchByID(ctx, newssettingsID)
}

func (lu *newssettingsUsecase) Fetch(c context.Context) ([]domain.NewsSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newssettingsRepository.Fetch(ctx)
}
