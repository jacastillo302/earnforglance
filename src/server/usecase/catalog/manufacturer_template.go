package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type manufacturertemplateUsecase struct {
	manufacturertemplateRepository domain.ManufacturerTemplateRepository
	contextTimeout                 time.Duration
}

func NewManufacturerTemplateUsecase(manufacturertemplateRepository domain.ManufacturerTemplateRepository, timeout time.Duration) domain.ManufacturerTemplateUsecase {
	return &manufacturertemplateUsecase{
		manufacturertemplateRepository: manufacturertemplateRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *manufacturertemplateUsecase) Create(c context.Context, manufacturertemplate *domain.ManufacturerTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.manufacturertemplateRepository.Create(ctx, manufacturertemplate)
}

func (tu *manufacturertemplateUsecase) Update(c context.Context, manufacturertemplate *domain.ManufacturerTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.manufacturertemplateRepository.Update(ctx, manufacturertemplate)
}

func (tu *manufacturertemplateUsecase) Delete(c context.Context, manufacturertemplate string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.manufacturertemplateRepository.Delete(ctx, manufacturertemplate)
}

func (lu *manufacturertemplateUsecase) FetchByID(c context.Context, manufacturertemplateID string) (domain.ManufacturerTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.manufacturertemplateRepository.FetchByID(ctx, manufacturertemplateID)
}

func (lu *manufacturertemplateUsecase) Fetch(c context.Context) ([]domain.ManufacturerTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.manufacturertemplateRepository.Fetch(ctx)
}
