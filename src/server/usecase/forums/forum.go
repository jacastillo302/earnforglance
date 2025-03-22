package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumUsecase struct {
	forumRepository domain.ForumRepository
	contextTimeout  time.Duration
}

func NewForumUsecase(forumRepository domain.ForumRepository, timeout time.Duration) domain.ForumUsecase {
	return &forumUsecase{
		forumRepository: forumRepository,
		contextTimeout:  timeout,
	}
}

func (tu *forumUsecase) Create(c context.Context, forum *domain.Forum) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumRepository.Create(ctx, forum)
}

func (tu *forumUsecase) Update(c context.Context, forum *domain.Forum) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumRepository.Update(ctx, forum)
}

func (tu *forumUsecase) Delete(c context.Context, forum string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumRepository.Delete(ctx, forum)
}

func (lu *forumUsecase) FetchByID(c context.Context, forumID string) (domain.Forum, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumRepository.FetchByID(ctx, forumID)
}

func (lu *forumUsecase) Fetch(c context.Context) ([]domain.Forum, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumRepository.Fetch(ctx)
}
