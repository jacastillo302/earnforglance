package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type categorytemplateUsecase struct {
	categorytemplateRepository domain.CategoryTemplateRepository
	contextTimeout             time.Duration
}

func NewCategoryTemplateUsecase(categorytemplateRepository domain.CategoryTemplateRepository, timeout time.Duration) domain.CategoryTemplateUsecase {
	return &categorytemplateUsecase{
		categorytemplateRepository: categorytemplateRepository,
		contextTimeout:             timeout,
	}
}

func (tu *categorytemplateUsecase) CreateMany(c context.Context, items []domain.CategoryTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.CreateMany(ctx, items)
}

func (tu *categorytemplateUsecase) Create(c context.Context, item *domain.CategoryTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.Create(ctx, item)
}

func (tu *categorytemplateUsecase) Update(c context.Context, item *domain.CategoryTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.Update(ctx, item)
}

func (tu *categorytemplateUsecase) Delete(c context.Context, item string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.Delete(ctx, item)
}

func (lu *categorytemplateUsecase) FetchByID(c context.Context, itemID string) (domain.CategoryTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.categorytemplateRepository.FetchByID(ctx, itemID)
}

func (lu *categorytemplateUsecase) Fetch(c context.Context) ([]domain.CategoryTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.categorytemplateRepository.Fetch(ctx)
}
