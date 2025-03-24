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

func (tu *categorytemplateUsecase) Create(c context.Context, affiliate *domain.CategoryTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.Create(ctx, affiliate)
}

func (tu *categorytemplateUsecase) Update(c context.Context, affiliate *domain.CategoryTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.Update(ctx, affiliate)
}

func (tu *categorytemplateUsecase) Delete(c context.Context, affiliate string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categorytemplateRepository.Delete(ctx, affiliate)
}

func (lu *categorytemplateUsecase) FetchByID(c context.Context, affiliateID string) (domain.CategoryTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.categorytemplateRepository.FetchByID(ctx, affiliateID)
}

func (lu *categorytemplateUsecase) Fetch(c context.Context) ([]domain.CategoryTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.categorytemplateRepository.Fetch(ctx)
}
