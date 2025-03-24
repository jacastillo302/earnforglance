package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumpostvoteUsecase struct {
	forumpostvoteRepository domain.ForumPostVoteRepository
	contextTimeout          time.Duration
}

func NewForumPostVoteUsecase(forumpostvoteRepository domain.ForumPostVoteRepository, timeout time.Duration) domain.ForumPostVoteUsecase {
	return &forumpostvoteUsecase{
		forumpostvoteRepository: forumpostvoteRepository,
		contextTimeout:          timeout,
	}
}

func (tu *forumpostvoteUsecase) CreateMany(c context.Context, items []domain.ForumPostVote) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostvoteRepository.CreateMany(ctx, items)
}

func (tu *forumpostvoteUsecase) Create(c context.Context, forumpostvote *domain.ForumPostVote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostvoteRepository.Create(ctx, forumpostvote)
}

func (tu *forumpostvoteUsecase) Update(c context.Context, forumpostvote *domain.ForumPostVote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostvoteRepository.Update(ctx, forumpostvote)
}

func (tu *forumpostvoteUsecase) Delete(c context.Context, forumpostvote string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumpostvoteRepository.Delete(ctx, forumpostvote)
}

func (lu *forumpostvoteUsecase) FetchByID(c context.Context, forumpostvoteID string) (domain.ForumPostVote, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumpostvoteRepository.FetchByID(ctx, forumpostvoteID)
}

func (lu *forumpostvoteUsecase) Fetch(c context.Context) ([]domain.ForumPostVote, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumpostvoteRepository.Fetch(ctx)
}
