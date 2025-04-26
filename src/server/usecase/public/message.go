package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type newsLetterUsecase struct {
	itemRepository domain.NewsLetterRepository
	contextTimeout time.Duration
}

func NewNewsLetterUsecase(itemRepository domain.NewsLetterRepository, timeout time.Duration) domain.NewsLetterRepository {
	return &newsLetterUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *newsLetterUsecase) NewsLetterSubscription(c context.Context, news domain.NewsLetterRequest, IpAdress string) (domain.NewsLetterResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.NewsLetterSubscription(ctx, news, IpAdress)
}

func (cu *newsLetterUsecase) NewsLetterUnSubscribe(c context.Context, news domain.NewsLetterRequest) (domain.NewsLetterResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.NewsLetterUnSubscribe(ctx, news)
}

func (cu *newsLetterUsecase) NewsLetterActivation(c context.Context, Guid string) (domain.NewsLetterResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.NewsLetterActivation(ctx, Guid)
}

func (cu *newsLetterUsecase) NewsLetterInactivate(c context.Context, Guid string) (domain.NewsLetterResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.NewsLetterInactivate(ctx, Guid)
}

func (cu *newsLetterUsecase) GetSlugs(c context.Context, record string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetSlugs(ctx, record)
}
