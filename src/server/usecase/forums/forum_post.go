package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumpostUsecase struct {
	forumpostRepository domain.ForumPostRepository
	contextTimeout      time.Duration
}

func NewForumPostUsecase(forumpostRepository domain.ForumPostRepository, timeout time.Duration) domain.ForumPostUsecase {
	return &forumpostUsecase{
		forumpostRepository: forumpostRepository,
		contextTimeout:      timeout,
	}
}

func (tu *forumpostUsecase) CreateMany(c context.Context, items []domain.ForumPost) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostRepository.CreateMany(ctx, items)
}

func (tu *forumpostUsecase) Create(c context.Context, forumpost *domain.ForumPost) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostRepository.Create(ctx, forumpost)
}

func (tu *forumpostUsecase) Update(c context.Context, forumpost *domain.ForumPost) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostRepository.Update(ctx, forumpost)
}

func (tu *forumpostUsecase) Delete(c context.Context, forumpost string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostRepository.Delete(ctx, forumpost)
}

func (lu *forumpostUsecase) FetchByID(c context.Context, forumpostID string) (domain.ForumPost, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumpostRepository.FetchByID(ctx, forumpostID)
}

func (lu *forumpostUsecase) Fetch(c context.Context) ([]domain.ForumPost, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumpostRepository.Fetch(ctx)
}
