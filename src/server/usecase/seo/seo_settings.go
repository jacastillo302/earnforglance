package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/seo"
)

type seosettingsUsecase struct {
	seosettingsRepository domain.SeoSettingsRepository
	contextTimeout        time.Duration
}

func NewSeoSettingsUsecase(seosettingsRepository domain.SeoSettingsRepository, timeout time.Duration) domain.SeoSettingsUsecase {
	return &seosettingsUsecase{
		seosettingsRepository: seosettingsRepository,
		contextTimeout:        timeout,
	}
}

func (tu *seosettingsUsecase) Create(c context.Context, seosettings *domain.SeoSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.seosettingsRepository.Create(ctx, seosettings)
}

func (tu *seosettingsUsecase) Update(c context.Context, seosettings *domain.SeoSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.seosettingsRepository.Update(ctx, seosettings)
}

func (tu *seosettingsUsecase) Delete(c context.Context, seosettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.seosettingsRepository.Delete(ctx, seosettings)
}

func (lu *seosettingsUsecase) FetchByID(c context.Context, seosettingsID string) (domain.SeoSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.seosettingsRepository.FetchByID(ctx, seosettingsID)
}

func (lu *seosettingsUsecase) Fetch(c context.Context) ([]domain.SeoSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.seosettingsRepository.Fetch(ctx)
}
