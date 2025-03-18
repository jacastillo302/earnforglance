package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type pdfsettingsUsecase struct {
	pdfsettingsRepository domain.PdfSettingsRepository
	contextTimeout        time.Duration
}

func NewPdfSettingsUsecase(pdfsettingsRepository domain.PdfSettingsRepository, timeout time.Duration) domain.PdfSettingsUsecase {
	return &pdfsettingsUsecase{
		pdfsettingsRepository: pdfsettingsRepository,
		contextTimeout:        timeout,
	}
}

func (tu *pdfsettingsUsecase) Create(c context.Context, pdfsettings *domain.PdfSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pdfsettingsRepository.Create(ctx, pdfsettings)
}

func (tu *pdfsettingsUsecase) Update(c context.Context, pdfsettings *domain.PdfSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pdfsettingsRepository.Update(ctx, pdfsettings)
}

func (tu *pdfsettingsUsecase) Delete(c context.Context, pdfsettings *domain.PdfSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pdfsettingsRepository.Delete(ctx, pdfsettings)
}

func (lu *pdfsettingsUsecase) FetchByID(c context.Context, pdfsettingsID string) (domain.PdfSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pdfsettingsRepository.FetchByID(ctx, pdfsettingsID)
}

func (lu *pdfsettingsUsecase) Fetch(c context.Context) ([]domain.PdfSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pdfsettingsRepository.Fetch(ctx)
}
