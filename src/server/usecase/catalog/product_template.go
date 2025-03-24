package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type producttemplateUsecase struct {
	producttemplateRepository domain.ProductTemplateRepository
	contextTimeout            time.Duration
}

func NewProductTemplateUsecase(producttemplateRepository domain.ProductTemplateRepository, timeout time.Duration) domain.ProductTemplateUsecase {
	return &producttemplateUsecase{
		producttemplateRepository: producttemplateRepository,
		contextTimeout:            timeout,
	}
}

func (tu *producttemplateUsecase) CreateMany(c context.Context, items []domain.ProductTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttemplateRepository.CreateMany(ctx, items)
}

func (tu *producttemplateUsecase) Create(c context.Context, producttemplate *domain.ProductTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttemplateRepository.Create(ctx, producttemplate)
}

func (tu *producttemplateUsecase) Update(c context.Context, producttemplate *domain.ProductTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttemplateRepository.Update(ctx, producttemplate)
}

func (tu *producttemplateUsecase) Delete(c context.Context, producttemplate string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producttemplateRepository.Delete(ctx, producttemplate)
}

func (lu *producttemplateUsecase) FetchByID(c context.Context, producttemplateID string) (domain.ProductTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.producttemplateRepository.FetchByID(ctx, producttemplateID)
}

func (lu *producttemplateUsecase) Fetch(c context.Context) ([]domain.ProductTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.producttemplateRepository.Fetch(ctx)
}
