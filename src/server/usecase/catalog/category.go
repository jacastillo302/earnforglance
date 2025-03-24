package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}

func (tu *categoryUsecase) CreateMany(c context.Context, items []domain.Category) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.categoryRepository.CreateMany(ctx, items)
}

func (cu *categoryUsecase) Create(c context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Create(ctx, category)
}

func (cu *categoryUsecase) Update(c context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Update(ctx, category)
}

func (cu *categoryUsecase) Delete(c context.Context, category string) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Delete(ctx, category)
}

func (cu *categoryUsecase) FetchByID(c context.Context, categoryID string) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.FetchByID(ctx, categoryID)
}

func (cu *categoryUsecase) Fetch(c context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Fetch(ctx)
}
