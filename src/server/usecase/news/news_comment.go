package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/news"
)

type newsCommentUsecase struct {
	newsCommentRepository domain.NewsCommentRepository
	contextTimeout        time.Duration
}

func NewNewsCommentUsecase(newsCommentRepository domain.NewsCommentRepository, timeout time.Duration) domain.NewsCommentUsecase {
	return &newsCommentUsecase{
		newsCommentRepository: newsCommentRepository,
		contextTimeout:        timeout,
	}
}

func (tu *newsCommentUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.NewsComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsCommentRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *newsCommentUsecase) Create(c context.Context, newsComment *domain.NewsComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsCommentRepository.Create(ctx, newsComment)
}

func (tu *newsCommentUsecase) Update(c context.Context, newsComment *domain.NewsComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsCommentRepository.Update(ctx, newsComment)
}

func (tu *newsCommentUsecase) Delete(c context.Context, newsComment string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsCommentRepository.Delete(ctx, newsComment)
}

func (lu *newsCommentUsecase) FetchByID(c context.Context, newsCommentID string) (domain.NewsComment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newsCommentRepository.FetchByID(ctx, newsCommentID)
}

func (lu *newsCommentUsecase) Fetch(c context.Context) ([]domain.NewsComment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newsCommentRepository.Fetch(ctx)
}
