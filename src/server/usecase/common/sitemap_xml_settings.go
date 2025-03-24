package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type sitemapxmlsettingsUsecase struct {
	sitemapxmlsettingsRepository domain.SitemapXmlSettingsRepository
	contextTimeout               time.Duration
}

func NewSitemapXmlSettingsUsecase(sitemapxmlsettingsRepository domain.SitemapXmlSettingsRepository, timeout time.Duration) domain.SitemapXmlSettingsUsecase {
	return &sitemapxmlsettingsUsecase{
		sitemapxmlsettingsRepository: sitemapxmlsettingsRepository,
		contextTimeout:               timeout,
	}
}

func (tu *sitemapxmlsettingsUsecase) CreateMany(c context.Context, items []domain.SitemapXmlSettings) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapxmlsettingsRepository.CreateMany(ctx, items)
}

func (tu *sitemapxmlsettingsUsecase) Create(c context.Context, sitemapxmlsettings *domain.SitemapXmlSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapxmlsettingsRepository.Create(ctx, sitemapxmlsettings)
}

func (tu *sitemapxmlsettingsUsecase) Update(c context.Context, sitemapxmlsettings *domain.SitemapXmlSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapxmlsettingsRepository.Update(ctx, sitemapxmlsettings)
}

func (tu *sitemapxmlsettingsUsecase) Delete(c context.Context, sitemapxmlsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.sitemapxmlsettingsRepository.Delete(ctx, sitemapxmlsettings)
}

func (lu *sitemapxmlsettingsUsecase) FetchByID(c context.Context, sitemapxmlsettingsID string) (domain.SitemapXmlSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.sitemapxmlsettingsRepository.FetchByID(ctx, sitemapxmlsettingsID)
}

func (lu *sitemapxmlsettingsUsecase) Fetch(c context.Context) ([]domain.SitemapXmlSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.sitemapxmlsettingsRepository.Fetch(ctx)
}
