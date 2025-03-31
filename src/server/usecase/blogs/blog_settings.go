package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/blogs"
)

type blogsettingsUsecase struct {
	blogsettingsRepository domain.BlogSettingsRepository
	contextTimeout         time.Duration
}

func NewBlogSettingsUsecase(blogsettingsRepository domain.BlogSettingsRepository, timeout time.Duration) domain.BlogSettingsUsecase {
	return &blogsettingsUsecase{
		blogsettingsRepository: blogsettingsRepository,
		contextTimeout:         timeout,
	}
}

func (tu *blogsettingsUsecase) CreateMany(c context.Context, items []domain.BlogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.CreateMany(ctx, items)
}

func (tu *blogsettingsUsecase) Create(c context.Context, item *domain.BlogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.Create(ctx, item)
}

func (tu *blogsettingsUsecase) Update(c context.Context, item *domain.BlogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.Update(ctx, item)
}

func (tu *blogsettingsUsecase) Delete(c context.Context, item string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.Delete(ctx, item)
}

func (lu *blogsettingsUsecase) FetchByID(c context.Context, itemID string) (domain.BlogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogsettingsRepository.FetchByID(ctx, itemID)
}

func (lu *blogsettingsUsecase) Fetch(c context.Context) ([]domain.BlogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogsettingsRepository.Fetch(ctx)
}
