package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumSubscriptionUsecase struct {
	forumSubscriptionRepository domain.ForumSubscriptionRepository
	contextTimeout              time.Duration
}

func NewForumSubscriptionUsecase(forumSubscriptionRepository domain.ForumSubscriptionRepository, timeout time.Duration) domain.ForumSubscriptionUsecase {
	return &forumSubscriptionUsecase{
		forumSubscriptionRepository: forumSubscriptionRepository,
		contextTimeout:              timeout,
	}
}

func (tu *forumSubscriptionUsecase) CreateMany(c context.Context, items []domain.ForumSubscription) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSubscriptionRepository.CreateMany(ctx, items)
}

func (tu *forumSubscriptionUsecase) Create(c context.Context, forumSubscription *domain.ForumSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSubscriptionRepository.Create(ctx, forumSubscription)
}

func (tu *forumSubscriptionUsecase) Update(c context.Context, forumSubscription *domain.ForumSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSubscriptionRepository.Update(ctx, forumSubscription)
}

func (tu *forumSubscriptionUsecase) Delete(c context.Context, forumSubscription string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSubscriptionRepository.Delete(ctx, forumSubscription)
}

func (lu *forumSubscriptionUsecase) FetchByID(c context.Context, forumSubscriptionID string) (domain.ForumSubscription, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumSubscriptionRepository.FetchByID(ctx, forumSubscriptionID)
}

func (lu *forumSubscriptionUsecase) Fetch(c context.Context) ([]domain.ForumSubscription, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumSubscriptionRepository.Fetch(ctx)
}
