package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

// Compare this snippet from src/server/usecase/common/sitemap_settings.go:
type sitemapsettingsUsecase struct {
	sitemapsettingsRepository domain.SitemapSettingsRepository
	contextTimeout            time.Duration
}

func NewSitemapSettingsUsecase(sitemapsettingsRepository domain.SitemapSettingsRepository, timeout time.Duration) domain.SitemapSettingsUsecase {
	return &sitemapsettingsUsecase{
		sitemapsettingsRepository: sitemapsettingsRepository,
		contextTimeout:            timeout,
	}
}

func (tu *sitemapsettingsUsecase) Create(c context.Context, sitemapsettings *domain.SitemapSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapsettingsRepository.Create(ctx, sitemapsettings)
}

func (tu *sitemapsettingsUsecase) Update(c context.Context, sitemapsettings *domain.SitemapSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapsettingsRepository.Update(ctx, sitemapsettings)
}

func (tu *sitemapsettingsUsecase) Delete(c context.Context, sitemapsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapsettingsRepository.Delete(ctx, sitemapsettings)
}

func (lu *sitemapsettingsUsecase) FetchByID(c context.Context, sitemapsettingsID string) (domain.SitemapSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.sitemapsettingsRepository.FetchByID(ctx, sitemapsettingsID)
}

func (lu *sitemapsettingsUsecase) Fetch(c context.Context) ([]domain.SitemapSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.sitemapsettingsRepository.Fetch(ctx)
}
