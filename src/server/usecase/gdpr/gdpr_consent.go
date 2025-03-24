package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/gdpr"
)

type gdprconsentUsecase struct {
	gdprconsentRepository domain.GdprConsentRepository
	contextTimeout        time.Duration
}

func NewGdprConsentUsecase(gdprconsentRepository domain.GdprConsentRepository, timeout time.Duration) domain.GdprConsentUsecase {
	return &gdprconsentUsecase{
		gdprconsentRepository: gdprconsentRepository,
		contextTimeout:        timeout,
	}
}

func (tu *gdprconsentUsecase) CreateMany(c context.Context, items []domain.GdprConsent) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.gdprconsentRepository.CreateMany(ctx, items)
}

func (tu *gdprconsentUsecase) Create(c context.Context, gdprconsent *domain.GdprConsent) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.gdprconsentRepository.Create(ctx, gdprconsent)
}

func (tu *gdprconsentUsecase) Update(c context.Context, gdprconsent *domain.GdprConsent) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.gdprconsentRepository.Update(ctx, gdprconsent)
}

func (tu *gdprconsentUsecase) Delete(c context.Context, gdprconsent string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.gdprconsentRepository.Delete(ctx, gdprconsent)
}

func (lu *gdprconsentUsecase) FetchByID(c context.Context, gdprconsentID string) (domain.GdprConsent, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.gdprconsentRepository.FetchByID(ctx, gdprconsentID)
}

func (lu *gdprconsentUsecase) Fetch(c context.Context) ([]domain.GdprConsent, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.gdprconsentRepository.Fetch(ctx)
}
