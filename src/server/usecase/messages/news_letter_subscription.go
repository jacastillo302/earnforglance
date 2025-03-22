package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type newsLetterSubscriptionUsecase struct {
	newsLetterSubscriptionRepository domain.NewsLetterSubscriptionRepository
	contextTimeout                   time.Duration
}

func NewNewsLetterSubscriptionUsecase(newsLetterSubscriptionRepository domain.NewsLetterSubscriptionRepository, timeout time.Duration) domain.NewsLetterSubscriptionUsecase {
	return &newsLetterSubscriptionUsecase{
		newsLetterSubscriptionRepository: newsLetterSubscriptionRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *newsLetterSubscriptionUsecase) Create(c context.Context, newsLetterSubscription *domain.NewsLetterSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsLetterSubscriptionRepository.Create(ctx, newsLetterSubscription)
}

func (tu *newsLetterSubscriptionUsecase) Update(c context.Context, newsLetterSubscription *domain.NewsLetterSubscription) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsLetterSubscriptionRepository.Update(ctx, newsLetterSubscription)
}

func (tu *newsLetterSubscriptionUsecase) Delete(c context.Context, newsLetterSubscription string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsLetterSubscriptionRepository.Delete(ctx, newsLetterSubscription)
}

func (lu *newsLetterSubscriptionUsecase) FetchByID(c context.Context, newsLetterSubscriptionID string) (domain.NewsLetterSubscription, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newsLetterSubscriptionRepository.FetchByID(ctx, newsLetterSubscriptionID)
}

func (lu *newsLetterSubscriptionUsecase) Fetch(c context.Context) ([]domain.NewsLetterSubscription, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.newsLetterSubscriptionRepository.Fetch(ctx)
}
