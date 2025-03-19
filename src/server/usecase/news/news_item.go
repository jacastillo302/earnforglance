package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/news"
)

type newsitemUsecase struct {
	newsitemRepository domain.NewsItemRepository
	contextTimeout     time.Duration
}

func NewNewsItemUsecase(newsitemRepository domain.NewsItemRepository, timeout time.Duration) domain.NewsItemUsecase {
	return &newsitemUsecase{
		newsitemRepository: newsitemRepository,
		contextTimeout:     timeout,
	}
}

func (tu *newsitemUsecase) Create(c context.Context, newsitem *domain.NewsItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsitemRepository.Create(ctx, newsitem)
}

func (tu *newsitemUsecase) Update(c context.Context, newsitem *domain.NewsItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsitemRepository.Update(ctx, newsitem)
}

func (tu *newsitemUsecase) Delete(c context.Context, newsitem *domain.NewsItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsitemRepository.Delete(ctx, newsitem)
}

func (lu *newsitemUsecase) FetchByID(c context.Context, newsitemID string) (domain.NewsItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newsitemRepository.FetchByID(ctx, newsitemID)
}

func (lu *newsitemUsecase) Fetch(c context.Context) ([]domain.NewsItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newsitemRepository.Fetch(ctx)
}
