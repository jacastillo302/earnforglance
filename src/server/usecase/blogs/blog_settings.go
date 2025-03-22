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

func (tu *blogsettingsUsecase) Create(c context.Context, affiliate *domain.BlogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.Create(ctx, affiliate)
}

func (tu *blogsettingsUsecase) Update(c context.Context, affiliate *domain.BlogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.Update(ctx, affiliate)
}

func (tu *blogsettingsUsecase) Delete(c context.Context, affiliate string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogsettingsRepository.Delete(ctx, affiliate)
}

func (lu *blogsettingsUsecase) FetchByID(c context.Context, affiliateID string) (domain.BlogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogsettingsRepository.FetchByID(ctx, affiliateID)
}

func (lu *blogsettingsUsecase) Fetch(c context.Context) ([]domain.BlogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogsettingsRepository.Fetch(ctx)
}
